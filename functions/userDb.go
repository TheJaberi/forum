package forum

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"forum"

	"github.com/gofrs/uuid"
	bcrypt "golang.org/x/crypto/bcrypt"
)

var (
	Session forum.Session
	Empty   forum.Session
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
	pass, _ := bcrypt.GenerateFromPassword([]byte(applicant.Password), 4)
	_, err = sqlStmt.Exec(applicant.Username, applicant.Email, pass, applicant.Type)
	if err != nil {
		return err
	}
	return nil
}

// Login
func UserDbLogin(email string, password string) (forum.Session, error) {
	if isUsernameExists(email) != nil {
		return Empty, UserEmailError
	}
	userdata := DB.QueryRow("SELECT user_id, user_name, user_pass, user_email, user_type FROM users where user_email = ?", email) // select gets the data from users table
	err := userdata.Scan(&LoggedUser.Userid, &LoggedUser.Username, &LoggedUser.Password, &LoggedUser.Email, &LoggedUser.Type)     // scan assigns the data of the row to variables
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
		return Empty, err
	}
	uuid, err := uuid.NewV4()
	if err != nil {
		return Empty, err
	}
	// XXX implement and return Cookie
	Session = forum.Session{
		Name:      "test",
		UserId:    LoggedUser.Userid,
		Email:     email,
		CreatedAt: time.Now(),
		Uuid:      uuid,
	}
	fmt.Println(Session)
	UpdatePosts()
	return Session, nil
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

func CheckCookies(r *http.Request) error {
	cookie, err := r.Cookie(Session.Name)
	if err != nil {
		return err
	}
	fmt.Println(cookie)
	// Get cookie value:
	if cookie.MaxAge < 0 {
		Session = Empty
		return errors.New("session expired")
	}
	return nil
}
