package forum

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"github.com/gofrs/uuid"
	bcrypt "golang.org/x/crypto/bcrypt"
)

var (
	// XXX Implement into db calls
	UserNameError     = errors.New("User Name error!")
	UserEmailError    = errors.New("User Email error!")
	UserPasswordError = errors.New("User Password error!")
	RegPasswordError =  errors.New("Password too weak!")
	UserExistsError   = errors.New("Email Already in Use!")
)

// Registeration
func UserDbRegisteration(applicant Applicant, db *sql.DB) error {
	if isEmailExists(applicant.Email) == nil {
		AllData.LoginErrorMsg = "Email Already in Use!"
		return UserExistsError
	}
	if !PasswordChecker(string(applicant.Password)){
		AllData.LoginErrorMsg = "Password too weak!\nmust be more than 6 characters"
		return RegPasswordError
	}
	if !NameChecker(applicant.Username){
		return UserNameError
	}
	sqlStmt, err := db.Prepare("INSERT INTO users (user_name, user_email, user_pass, user_type) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(applicant.Password), 4)
	if err != nil {
		AllData.LoginErrorMsg = "Password too weak!"
		return RegPasswordError
	}
	_, err = sqlStmt.Exec(applicant.Username, applicant.Email, pass, applicant.Type)
	if err != nil {
		return err
	}
	return nil
}

// Login
func UserDbLogin(email string, password string) error {
	if isEmailExists(email) != nil {
		AllData.LoginErrorMsg = "User Email error!"
		return UserEmailError
	}
	userdata := DB.QueryRow("SELECT user_id, user_name, user_pass, user_email, user_type FROM users where user_email = ?", email) // select gets the data from users table	
	err := userdata.Scan(&LoggedUser.Userid, &LoggedUser.Username, &LoggedUser.Password, &LoggedUser.Email, &LoggedUser.Type) // scan assigns the data of the row to variables
	fmt.Println(LoggedUser.Username)
	if err != nil || LoggedUser.Username == ""{
		return err
			} else {
				err2 := bcrypt.CompareHashAndPassword([]byte(LoggedUser.Password), []byte(password))
				if err2 != nil {
					AllData.LoginErrorMsg = "User Password error!"
					return err2
				}
		LoggedUser.Registered = true
		AllData.IsLogged = true
		AllData.LoggedUser = LoggedUser
		AllData.LoggedUserID = LoggedUser.Userid
		if LoggedUser.Type == "admin" {
			AllData.TypeAdmin = true
		}
	}
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	// XXX implement and return Cookie
	LiveSession = Session{
		Id: LoggedUser.Userid,
		Email:     email,
		CreatedAt: time.Now(),
		Uuid:      uuid,
	}
	fmt.Println(LiveSession)
	UpdatePosts()	
	return nil
}

func isEmailExists(applicantEmail string) error {
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
