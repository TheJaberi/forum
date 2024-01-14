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
	Id        int
	Uuid      uuid.UUID
	Email     string
	UserId    int
	CreatedAt time.Time
}
