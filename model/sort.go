package forum

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
