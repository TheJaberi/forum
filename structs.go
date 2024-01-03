package forum

import (
	"time"

	"github.com/gofrs/uuid"
)

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

type Comment struct {
	UserName      string
	UserId        int
	Body          string
	Count_Likes   int
	Count_Dislike int
}
