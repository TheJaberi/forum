package forum

import (
	"strings"
)

func SortPosts(sortby string) {
	if sortby == "oldest" {
		AllData.AllPosts = AllPosts
	} else if sortby == "mostliked" {
		AllData.AllPosts = SortByLike(AllPosts)
	} else if sortby == "mostdisliked" {
		AllData.AllPosts = SortByDislike(AllPosts)
	} else if sortby == "mostcommentedon" {
		AllData.AllPosts = SortByComment(AllPosts)
	}
}

func RSort(list []Post) []Post {
	var arrAllPosts []Post
	for i := len(list) - 1; i >= 0; i-- {
		arrAllPosts = append(arrAllPosts, list[i])
	}
	return arrAllPosts
}

func SortByLike(list []Post) []Post {
	var arrAllPosts []Post
	for i := 0; i >= 0; i++ {
		for j := 0; j < len(list); j++ {
			if list[j].Likes == i {
				arrAllPosts = append(arrAllPosts, list[j])
			}
		}
		if len(arrAllPosts) >= len(list) {
			break
		}
	}
	return RSort(arrAllPosts)
}

func SortByDislike(list []Post) []Post {
	var arrAllPosts []Post
	for i := 0; i >= 0; i++ {
		for j := 0; j < len(list); j++ {
			if list[j].Dislikes == i {
				arrAllPosts = append(arrAllPosts, list[j])
			}
		}
		if len(arrAllPosts) >= len(list) {
			break
		}
	}
	return RSort(arrAllPosts)
}

func SortByComment(list []Post) []Post {
	var arrAllPosts []Post
	for i := 0; i >= 0; i++ {
		for j := 0; j < len(list); j++ {
			if len(list[j].Comments) == i {
				arrAllPosts = append(arrAllPosts, list[j])
			}
		}
		if len(arrAllPosts) >= len(list) {
			break
		}
	}
	return RSort(arrAllPosts)
}
func AdjustText(text string) string {
	arrText := strings.Split(text, " ")
	finaltext := ""
	for i := 0; i < len(arrText); i++ {
		if len(arrText[i]) > 25 {
			for j := 25; j < len(arrText[i]); {
				arrText[i] = arrText[i][:j] + " " + arrText[i][j:]
				j += 25
			}
			finaltext = finaltext + arrText[i]
		} else {
			finaltext = finaltext + " " + arrText[i]
		}
	}
	return finaltext
}
