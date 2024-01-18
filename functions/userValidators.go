package forum

import "regexp"

func NameChecker(name string) bool {
	if len(name) < 3 || len(name) > 14 {
		AllData.LoginErrorMsg = "Username must be between 3 and 13 characters"
		return false
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !re.MatchString(name){
		AllData.LoginErrorMsg = "Username must contain numbers and letters only"
		return false
	} else {
return true
	}
}

func PasswordChecker(pass string) bool {
	if len(pass) < 6 && len(pass) > 25 {
		return false
	}
	for _, r := range pass {
		if r < 32 && r > 126 {
			return false
		}
	}
	// for i := 0; i < len(pass)-2; i++ {
	// 	if pass[i] == pass[i+1] && pass[i+1] == pass[i+2] {
	// 		return false
	// 	}
	// }
	return true
}

func EmailChecker(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
