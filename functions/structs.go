package forum

import "database/sql"

var (
	Database      *sql.DB
	LoggedUser    User
	ErrorMsg      string
	AllPosts      []Post
	AllCategories []Category
	AllData       Data
)

type Data struct {
	AllPosts      []Post
	AllCategories []Category
}
type Category struct {
	CategoryName string
	CategoryID   int
}
type User struct {
	Userid     int
	Username   string
	Password   string
	Email      string
	Registered bool
	Type       string
}

var ErrResponse struct {
	StatusCode bool
	ErrorMsg   string
}

type Post struct {
	Title    string
	Body     string
	PostID   int
	UserID   int
	Username string
	Category []Category
	Likes int
	Dislikes int
	Userlike bool
	UserDislike bool
}
