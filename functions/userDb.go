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
func UserDbLogin(email string, password string) error {
	if isUsernameExists(email) != nil {
		return UserEmailError
	}
	userdata := DB.QueryRow("SELECT user_id, user_name, user_pass, user_email, user_type FROM users where user_email = ?", email) // select gets the data from users table
	err := userdata.Scan(&LoggedUser.Userid, &LoggedUser.Username, &LoggedUser.Password, &LoggedUser.Email, &LoggedUser.Type) // scan assigns the data of the row to variables
	if err != nil {
		fmt.Println(err)
			} else {
		LoggedUser.Registered = true
		AllData.IsLogged = true
		AllData.LoggedUser = LoggedUser
		AllData.LoggedUserID = LoggedUser.Userid
		if LoggedUser.Type == "admin" {
			AllData.TypeAdmin = true
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(LoggedUser.Password), []byte(password))
	if err != nil {
		return err
	}
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	// XXX implement and return Cookie
	session := forum.Session{
		Id: LoggedUser.Userid,
		Email:     email,
		CreatedAt: time.Now(),
		Uuid:      uuid,
	}
	fmt.Println(session)
	UpdatePosts()	
	return nil
}

func isUsernameExists(applicantEmail string) error {
	sqlStmt := `SELECT EXISTS (SELECT 1 FROM users WHERE user_email = ?)`
	var exists bool
	err := DB.QueryRow(sqlStmt, applicantEmail).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return UserExistsError
	}
	return nil
}
