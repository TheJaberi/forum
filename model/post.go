package forum

import (
	"context"
	"errors"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func CreatePost(title string, body string, postCategories []int) error {
	var postData = Post{
		Title:    title,
		Body:     body,
		UserID:   LoggedUser.Userid,
		Username: LoggedUser.Username,
	}
	id := createPostDb(postData)
	err := assignPostCategoryDb(id, postCategories)
	if err != nil {
		return err
	}
	newPost, err := getPost(id)
	if err != nil {
		return err
	}
	AllData.AllPosts = append(AllPosts, newPost)
	return nil
}

func createPostDb(post Post) int64 {
	query := "INSERT INTO `posts` (`Title`, `body`, `user_id`) VALUES (?, ?, ?)"
	rowData, err := DB.ExecContext(context.Background(), query, post.Title, post.Body, post.UserID)
	if err != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err)
	}
	postID, err := rowData.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return postID
}

func assignPostCategoryDb(postID int64, postCategories []int) error {
	for _, category := range postCategories {
		queryCategory := "INSERT INTO `Post2Category` (`post_id`, `category_id`) VALUES (?, ?)"
		_, err := DB.ExecContext(context.Background(), queryCategory, postID, category)
		if err != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
			log.Fatal(err)
			return err
		}
		/*
			for j := 0; j < len(AllCategories); j++ {
				if AllCategories[j].CategoryID == postCategories[i] {
					postData.Category = append(postData.Category, AllCategories[j])
					break
				}
			}
		*/
	}
	err := GetCategories()
	if err != nil {
		return err
	}
	return nil
}

func getPost(id int64) (Post, error) {
	row := DB.QueryRow("SELECT id, Title, body, user_id, time_created from posts WHERE id=?", id)
	var p Post
	err := row.Scan(&p.PostID, &p.Title, &p.Body, &p.UserID, &p.TimeCreated)
	if err != nil {
		log.Printf("Error Getting Post")
		return p, err
	}
	p.TimeCreated = strings.Replace(p.TimeCreated, "T", " ", -1)
	p.TimeCreated = strings.Replace(p.TimeCreated, "Z", " ", -1)
	return p, nil
}

// this belongs in interactions?
func getUserPostInteractions() error {
	if LoggedUser.Registered { // FIXME Possibly redundent
		for i := range AllPosts {
			var interaction int
			postData := DB.QueryRow("SELECT interaction from Interaction where post_id = ? AND user_id = ?", i+1, LoggedUser.Userid)
			errpost := postData.Scan(&interaction)
			if errpost != nil {
				// FIXME what is the case for this?
				// continue
				return errors.New("Error scanning user post interactions")
			}
			if interaction == 1 {
				AllPosts[i].Userlike = true
			} else {
				AllPosts[i].UserDislike = true
			}
		}
	} else {
		return errors.New("User must log in to get post interactions")
	}
	return nil
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
		p, err = getPostDetails(p)
		AllPosts = append(AllPosts, p)
	}
	if LoggedUser.Registered {
		err := getUserPostInteractions()
		if err != nil {
			return err
		}
	}
	AllData.AllPosts = AllPosts
	return nil
}

func getPostDetails(p Post) (Post, error) {
	// USERNAME
	err := getPostUsername(&p)
	if err != nil {
		log.Println(err.Error())
		return p, err
	}
	// CATEGORIES
	err = getPostCategories(&p)
	if err != nil {
		log.Println(err.Error())
		return p, err
	}
	// COMMENTS
	err = getPostComments(&p)
	if err != nil {
		log.Println(err.Error())
		return p, err
	}
	// INTERACTIONS
	err = getPostLikes(&p)
	if err != nil {
		log.Println(err.Error())
		return p, err
	}
	err = getPostDislikes(&p)
	if err != nil {
		log.Println(err.Error())
		return p, err
	}
	return p, nil
}

func getPostUsername(p *Post) error {
	userData := DB.QueryRow("Select user_name from users where user_id = ?", p.PostID)
	err := userData.Scan(&p.Username)
	if err != nil {
		return errors.New("User Scan Error:" + err.Error())
	}
	return nil
}

func getPostCategories(p *Post) error {
	categoryData, err := DB.Query("Select category_id from Post2Category where post_id = ?", p.PostID)
	if err != nil {
		return errors.New("Category Query Error:" + err.Error())
	}
	defer categoryData.Close()
	for categoryData.Next() {
		var categoryID int
		err := categoryData.Scan(&categoryID)
		if err != nil {
			return errors.New("Category Scan Error:" + err.Error())
		}
		for i := range AllCategories {
			if categoryID == AllCategories[i].CategoryID {
				p.Category = append(p.Category, AllCategories[i])
				break
			}
		}
	}
	return nil
}

func getPostComments(p *Post) error {
	commentData, err := DB.Query("Select comment_id, body, user_id, time_created from comments where post_id = ?", p.PostID)
	if err != nil {
		return errors.New("Comment Query Error:" + err.Error())
	}
	defer commentData.Close()
	for commentData.Next() {
		var c Comment
		err := commentData.Scan(&c.Comment_id, &c.Body, &c.User_id, &c.TimeCreated)
		if err != nil {
			return errors.New("Comment Scan Error:" + err.Error())
		}
		c.TimeCreated = strings.Replace(c.TimeCreated, "T", " ", -1)
		c.TimeCreated = strings.Replace(c.TimeCreated, "Z", " ", -1)
		err = GetCommentUsername(&c)
		if err != nil {
			return err
		}
		p.Comments = append(p.Comments, c)
	}
	p.NumbOfComments = len(p.Comments)
	return nil
}

func getPostLikes(p *Post) error {
	likedata := DB.QueryRow("SELECT COUNT(user_id) FROM Interaction where post_id = ? AND interaction = ?", p.PostID, 1)
	err := likedata.Scan(&p.Likes)
	if err != nil {
		return errors.New("Post Interaction (Likes) Scan Error:" + err.Error())
	}
	return nil
}

func getPostDislikes(p *Post) error {
	dislikedata := DB.QueryRow("SELECT COUNT(user_id) FROM Interaction where post_id = ? AND interaction = ?", p.PostID, 0)
	err := dislikedata.Scan(&p.Dislikes)
	if err != nil {
		return errors.New("Post Interaction (Dislikes) Scan Error:" + err.Error())
	}
	return nil
}
