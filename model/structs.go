package forum

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
)

var (
	Database      *sql.DB
	LoggedUser    User
	ErrorMsg      string
	AllCategories []Category
	AllData       Data
	AllPosts      []Post
	LiveSession   Session
	EmptySession  Session
	Empty         User
	LoginError2   bool
)

type Data struct {
	AllPosts      []Post
	AllCategories []Category
	Postpage      Post
	LoggedUser    User
	CategoryCheck bool
	IsLogged      bool
	LoggedUserID  int
	TypeAdmin     bool
	LoginError    bool
	LoginErrorMsg string
}
type Category struct {
	CategoryName string
	CategoryID   int
}
type User struct {
	Userid     int
	Username   string
	Password   string // TODO is this necessary? Can we hash prior to assigning?
	Email      string
	Registered bool
	Type       string
}

var ErrResponse struct {
	StatusCode bool
	ErrorMsg   string
}

type Post struct {
	Title          string
	Body           string
	PostID         int
	UserID         int
	Username       string
	Category       []Category
	Likes          int
	Dislikes       int
	Userlike       bool
	UserDislike    bool
	LoggedUser     bool
	TimeCreated    string
	Comments       []Comment
	NumbOfComments int
	Image          string
}

type Comment struct {
	Body               string
	Post_id            int
	User_id            int
	CommentUsername    string
	TimeCreated        string
	Likes              int
	Dislikes           int
	CommentUserlike    bool
	CommentUserDislike bool
	CommentLoggedUser  bool
	Comment_id         int
}

type Applicant struct {
	Username string
	Email    string
	Password []byte
	Type     string
}

type Session struct {
	Name      string
	Uuid      uuid.UUID
	Email     string
	UserId    int
	CreatedAt time.Time
}
