package forum

import "errors"

// User Error Messages
var (
	UserNameError       = errors.New("User Name error!")
	UserNameLengthError = errors.New("Username must be between 3 and 13 characters")
	UserNameCharError   = errors.New("Username must contain numbers and letters only")
	UserEmailError      = errors.New("Email does not exist!")
	UserPasswordError   = errors.New("User Password error!")
	RegPasswordError    = errors.New("Password too weak!\nmust be more than 6 characters")
	UserExistsError     = errors.New("Email Already in Use!")
	SessionExpired      = errors.New("Session Expired")
	EmailFormatError    = errors.New("Wrong Email Format!")
)
