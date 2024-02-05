package forum

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	bcrypt "golang.org/x/crypto/bcrypt"
)

// Receive new user data, validate and insert in user table
func UserRegisteration(applicant Applicant, db *sql.DB) error {
	err := RegisterValidator(applicant)
	if err != nil {
		AllData.LoginErrorMsg = err.Error()
		return err
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(applicant.Password), 4)
	if err != nil {
		AllData.LoginErrorMsg = UserPasswordError.Error() // Using simplified error for user
		return err
	}
	err = UserInsertDb(applicant, db, pass)
	if err != nil {
		return err
	}
	return nil
}

// Receive login credentials, validate and respond with a session cookie
func UserLogin(email string, password string) (Session, error) {
	// Validate User Existance
	if UserExistsDb(email) != nil {
		AllData.LoginErrorMsg = UserEmailError.Error()
		return EmptySession, UserEmailError
	}
	// Validates Entered Password
	err := bcrypt.CompareHashAndPassword([]byte(LoggedUser.Password), []byte(password))
	if err != nil {
		AllData.LoginErrorMsg = err.Error()
		return EmptySession, err
	}
	// Retrieves User Data
	err = UserRetrieveDb(email, password)
	if err != nil {
		return EmptySession, err
	}

	LoggedUser.Registered = true
	AllData.IsLogged = true
	AllData.LoggedUser = LoggedUser
	AllData.LoggedUserID = LoggedUser.Userid
	if LoggedUser.Type == "admin" {
		AllData.TypeAdmin = true
	}
	session, err := CreateCookies()
	if err != nil {
		return EmptySession, err
	}
	//UpdatePosts()
	return session, nil
}

// Create a session cookie from the user data from the global struct
func CreateCookies() (Session, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return EmptySession, err
	}
	LiveSession = Session{
		Name:      "myCookies",
		UserId:    LoggedUser.Userid,
		Email:     LoggedUser.Email,
		CreatedAt: time.Now(),
		Uuid:      uuid,
	}
	return LiveSession, nil
}

// Validate session remaining time
func CheckCookies(r *http.Request) error {
	cookie, err := r.Cookie(LiveSession.Name)
	if err != nil {
		return err
	}
	// Get cookie value:
	if cookie.MaxAge < 0 {
		LiveSession = EmptySession
		return SessionExpired
	}
	return nil
}
