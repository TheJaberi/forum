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
func UserLogin(email string, password string) (*http.Cookie, error) {
	var EmptyCookie *http.Cookie
	// Validate User Existance
	if UserExistsDb(email) != nil {
		AllData.LoginErrorMsg = UserEmailError.Error()
		return EmptyCookie, UserEmailError
	}
	// Retrieves User Data
	err := UserRetrieveDb(email, password)
	if err != nil {
		return EmptyCookie, err
	}
	// Validates Entered Password
	err = bcrypt.CompareHashAndPassword([]byte(LoggedUser.Password), []byte(password))
	if err != nil {
		AllData.LoginErrorMsg = err.Error()
		return EmptyCookie, err
	}

	LoggedUser.Registered = true
	AllData.IsLogged = true
	AllData.LoggedUser = LoggedUser
	AllData.LoggedUserID = LoggedUser.Userid

	if LoggedUser.Type == "admin" {
		AllData.TypeAdmin = true
	}
	session, err := CreateSession()
	if err != nil {
		return EmptyCookie, err
	}
	cookie := CreateCookie(session)

	err = GetUserPostInteractions()
	if err != nil {
		return EmptyCookie, err
	}
	return cookie, nil
}

// Create a cookie using the user session
func CreateCookie(s Session) *http.Cookie {
	c := &http.Cookie{
		Name:     s.Name,
		Value:    s.Uuid.String(),
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	}
	return c
}

// Create a session from the user data from the global struct
func CreateSession() (Session, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return EmptySession, err
	}
	LiveSession = Session{
		Name:      LoggedUser.Username,
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
