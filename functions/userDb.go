package forum

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"forum"

	"github.com/gofrs/uuid"
	bcrypt "golang.org/x/crypto/bcrypt"
)

var (
	// XXX Implement into db calls
	UserNameError     = errors.New("User Name error!")
	UserEmailError    = errors.New("User Email error!")
	UserPasswordError = errors.New("User Password error!")
	UserExistsError   = errors.New("Email Already in Use!")
)

// Registeration
func UserDbRegisteration(applicant forum.Applicant, db *sql.DB) error {
	sqlStmt, err := db.Prepare("INSERT INTO users (user_name, user_email, user_pass, user_type) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = sqlStmt.Exec(applicant.Username, applicant.Email, applicant.Password, applicant.Type)
	if err != nil {
		return err
	}
	return nil
}

// Login
func UserDbLogin(credentials forum.Login, db *sql.DB) error {
	if isUsernameExists(db, credentials.Email) != nil {
		return UserEmailError
	}
	sqlStmt, err := db.Prepare("SELECT user_pass FROM users WHERE user_email = ?")
	if err != nil {
		return err
	}
	var currentPass []byte
	err = sqlStmt.QueryRow(credentials.Email).Scan(&currentPass)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(currentPass, []byte(credentials.Password))
	if err != nil {
		return err
	}
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	// XXX implement and return Cookie
	session := forum.Session{
		Email:     credentials.Email,
		CreatedAt: time.Now(),
		Uuid:      uuid,
	}
	fmt.Println(session)
	return nil
}

func isUsernameExists(db *sql.DB, applicantEmail string) error {
	sqlStmt := `SELECT EXISTS (SELECT 1 FROM users WHERE user_email = ?)`
	var exists bool
	err := db.QueryRow(sqlStmt, applicantEmail).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return UserExistsError
	}
	return nil
}
