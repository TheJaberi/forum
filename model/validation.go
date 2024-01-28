package forum

import "regexp"

// Validate all registeration requirements
func RegisterValidator(applicant Applicant) error {
	if UserExistsDb(applicant.Email) == nil {
		return UserExistsError
	}
	err := emailChecker(applicant.Email)
	if err != nil {
		return err
	}
	err = passwordChecker(string(applicant.Password))
	if err != nil {
		return err
	}
	err = nameChecker(applicant.Username)
	if err != nil {
		return UserNameError
	}
	return nil
}

// Validate email format
func emailChecker(email string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return EmailFormatError
	}
	return nil
}

// Validate password length and characters
func passwordChecker(pass string) error {
	if len(pass) < 6 && len(pass) > 25 {
		return RegPasswordError
	}
	for _, r := range pass {
		if r < 32 && r > 126 {
			return RegPasswordError
		}
	}
	return nil
}

// Validate username length and characters
func nameChecker(name string) error {
	if len(name) < 3 || len(name) > 14 {
		return UserNameLengthError
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !re.MatchString(name) {
		return UserNameCharError
	}
	return nil
}
