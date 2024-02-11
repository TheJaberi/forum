package forum

import (
	"log"
	"strconv"
)

func FilterByCategory(categoryID string) error {
	var filteredPosts []Post
	GetCategories()
	GetPosts()
	category, err := strconv.Atoi(categoryID)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	for i := 0; i < len(AllPosts); i++ {
		for j := 0; j < len(AllPosts[i].Category); j++ {
			if category == AllPosts[i].Category[j].CategoryID {
				filteredPosts = append(filteredPosts, AllPosts[i])
				break
			}
		}
	}
	AllData.AllPosts = RSort(filteredPosts)
	AllData.CategoryCheck = false
	AllData.LoggedUser = LoggedUser
	return nil
}

func FilterUserData(userID, path string) error {
	var filteredPosts []Post
	user_id, err := strconv.Atoi(userID)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	for i := 0; i < len(AllPosts); i++ {
		if path == "/myposts/" && AllPosts[i].UserID == user_id {
			filteredPosts = append(filteredPosts, AllPosts[i])
		}
		if path == "/mylikes/" && AllPosts[i].Userlike {
			filteredPosts = append(filteredPosts, AllPosts[i])
		}
		if path == "/mydislikes/" && AllPosts[i].UserDislike {
			filteredPosts = append(filteredPosts, AllPosts[i])
		}
	}
	AllData.AllPosts = RSort(filteredPosts)
	AllData.CategoryCheck = false
	return nil
}
