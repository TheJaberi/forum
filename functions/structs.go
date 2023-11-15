package forum

import "database/sql"

var (
	Database   *sql.DB
	LoggedUser User
	ErrorMsg   string
	AllPosts   []Post
	AllCategories []string
	AllData Data
)

type Data struct {
	AllPosts []Post
	AllCategories []string
}
type Category struct {
	
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
	Title   string
	Body    string
	PostID int
	UserID int
	Username string
	Category string
}
