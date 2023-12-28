package forum

import (
	"time"

	"github.com/gofrs/uuid"
)

type Login struct {
	Email    string
	Password string
}

type Applicant struct {
	Username string
	Email    string
	Password []byte
	Type     string
}

type Session struct {
	Id        int
	Uuid      uuid.UUID
	Email     string
	UserId    int
	CreatedAt time.Time
}

type User struct {
	Name     string
	Type     string
	Username string
	Email    string
	Password []byte
}

type Post struct {
	UserName      string
	UserId        int
	Title         string
	Img           string
	Body          string
	Count_Likes   int
	Count_Dislike int
	Comments      Comment
}

type Comment struct {
	UserName      string
	UserId        int
	Body          string
	Count_Likes   int
	Count_Dislike int
}
