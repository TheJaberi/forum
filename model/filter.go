package forum

import (
	"log"
	"strconv"
)

func FilterByCategory(c string) error {
	var filteredPosts []Post
	GetCategories()
	GetPosts()
	category, err := strconv.Atoi(c)
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
