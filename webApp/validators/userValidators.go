package forum

import "regexp"

func NameChecker(name string) bool {
	if len(name) < 2 && len(name) > 8 {
		return false
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return re.MatchString(name)
}

func PasswordChecker(pass string) bool {
	if len(pass) < 8 && len(pass) > 25 {
		return false
	}
	for _, r := range pass {
		if r < 32 && r > 126 {
			return false
		}
	}
	for i := 0; i < len(pass)-2; i++ {
		if pass[i] == pass[i+1] && pass[i+1] == pass[i+2] {
			return false
		}
	}
	return true
}

func EmailChecker(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
