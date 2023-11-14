package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	bcrypt "golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data Register
	data.User_name = r.PostFormValue("user_name")
	if !NameChecker(data.User_name) {
		fmt.Println("User Name error!")
	}
	data.User_email = r.PostFormValue("user_email")
	if !EmailChecker(data.User_email) {
		fmt.Println("User Email error!")
	}
	pass := r.PostFormValue("password")
	if !PasswordChecker(pass) {
		fmt.Println("User Password error!")
	}
	data.User_pass, _ = bcrypt.GenerateFromPassword([]byte(pass), 4)
	data.User_type = "member"
	println(data.User_name, data.User_email, data.User_pass, data.User_type)
	sqlStmt, err := db.Prepare("INSERT INTO users (user_name, user_email, user_pass, user_type) VALUES (?, ?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sqlStmt.Exec(data.User_name, data.User_email, data.User_pass, data.User_type)
	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "test.html", nil)
}

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
