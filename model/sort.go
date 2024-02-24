package forum

// sorts the post depending on the link clicked in the frontend
func SortPosts(sortby string) {
	if sortby == "oldest" {
		AllData.AllPosts = AllPosts
	} else if sortby == "mostliked" {
		AllData.AllPosts = SortByLike(AllPosts)
	} else if sortby == "mostdisliked" {
		AllData.AllPosts = SortByDislike(AllPosts)
	} else if sortby == "mostcommentedon" {
		AllData.AllPosts = SortByComment(AllPosts)
	} else {
		if sortby != "" {
			FilterUserData(sortby)
		}
	}
}

// sorts posts in reverse
func RSort(list []Post) []Post {
	var arrAllPosts []Post
	for i := len(list) - 1; i >= 0; i-- {
		arrAllPosts = append(arrAllPosts, list[i])
	}
	return arrAllPosts
}

// sorts posts depending on number of likes
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

// sorts posts depending on number of dislikes
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

// sorts posts depending on number of comments
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

func RemoveSpaces(text string) string {
	var final []byte
	var wordstart int
	strbyte := []byte(text)
	for i := 0; i < len(text); i++ {
		if strbyte[i] != 32 {
			wordstart = i
			break
		}
	}
	if wordstart == 0 && strbyte[0] == 32 {
		return ""
	}
	for j := wordstart; j < len(text); j++ {
		final = append(final, strbyte[j])
	}
	return string(final)
}
