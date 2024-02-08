package forum

import (
	"context"
	"errors"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// TABLE: posts

func CreatePost(title string, body string, postCategories []int) error {
	var postData = Post{
		Title:    title,
		Body:     body,
		UserID:   LoggedUser.Userid,
		Username: LoggedUser.Username,
	}
	id := CreatePostDb(postData)
	err := AssignPostCategoryDb(id, postCategories)
	if err != nil {
		return err
	}
	newPost, err := GetPost(id)
	if err != nil {
		return err
	}
	AllData.AllPosts = append(AllPosts, newPost)
	return nil
}

func CreatePostDb(post Post) int {
	query := "INSERT INTO `posts` (`Title`, `body`, `user_id`) VALUES (?, ?, ?)"
	rowData, err := DB.ExecContext(context.Background(), query, post.Title, post.Body, post.UserID)
	if err != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err)
	}
	postID, err := rowData.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return int(postID)
}

func GetPost(id int) (Post, error) {
	row := DB.QueryRow("SELECT id, Title, body, user_id, time_created from posts WHERE id=?", id)
	var p Post
	err := row.Scan(&p.PostID, &p.Title, &p.Body, &p.UserID, &p.TimeCreated)
	if err != nil {
		log.Println(id)
		log.Printf("Error Getting Post")
		return p, err
	}
	p.TimeCreated = strings.Replace(p.TimeCreated, "T", " ", -1)
	p.TimeCreated = strings.Replace(p.TimeCreated, "Z", " ", -1)
	p, err = GetPostDetails(p)
	if err != nil {
		log.Println(errors.New("Post Scan Error: " + err.Error()))
		return p, err
	}
	return p, nil
}

func GetPosts() error {
	AllPosts = nil // FIXME Why do we clear it all the time?
	postData, errpost := DB.Query("Select id, Title, body, user_id, time_created from posts")
	if errpost != nil {
		log.Fatal(errpost)
	}
	defer postData.Close()
	for postData.Next() {
		var p Post
		err := postData.Scan(&p.PostID, &p.Title, &p.Body, &p.UserID, &p.TimeCreated)
		if err != nil {
			log.Println(errors.New("Post Scan Error: " + err.Error()))
			return err
		}
		p.TimeCreated = strings.Replace(p.TimeCreated, "T", " ", -1)
		p.TimeCreated = strings.Replace(p.TimeCreated, "Z", " ", -1)
		p, err = GetPostDetails(p)
		AllPosts = append(AllPosts, p)
	}
	if LoggedUser.Registered {
		err := GetUserPostInteractions()
		if err != nil {
			return err
		}
	}
	AllData.AllPosts = AllPosts
	return nil
}

func GetPostDetails(p Post) (Post, error) {
	// USERNAME
	err := GetPostUsername(&p)
	if err != nil {
		log.Println(err.Error())
		return p, err
	}
	// CATEGORIES
	err = GetPostCategories(&p)
	if err != nil {
		log.Println(err.Error())
		return p, err
	}
	// COMMENTS
	err = GetPostComments(&p)
	if err != nil {
		log.Println(err.Error())
		return p, err
	}
	// INTERACTIONS
	err = GetPostLikes(&p)
	if err != nil {
		log.Println(err.Error())
		return p, err
	}
	err = GetPostDislikes(&p)
	if err != nil {
		log.Println(err.Error())
		return p, err
	}
	return p, nil
}
