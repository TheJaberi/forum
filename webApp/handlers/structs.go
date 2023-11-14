package forum

import (
	"time"

	"github.com/gofrs/uuid"
)

type Login struct {
	User_email string
	User_pass  string
}

type Register struct {
	User_name  string
	User_email string
	User_pass  []byte
	User_type  string
}

type Session struct {
	Id        int
	Uuid      uuid.UUID
	Email     string
	UserId    int
	CreatedAt time.Time
}
