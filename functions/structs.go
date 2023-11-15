package forum

import "database/sql"

var (
	Database   *sql.DB
	LoggedUser User
	ErrorMsg   string
	AllPosts   []Post
)

type User struct {
	Userid     int
	Username   string
	Password   string
	Email      string
	Registered bool
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
