package forum

import (
	"database/sql"
	"forum"
	UserDb "forum/database"
	UserV "forum/webApp/validators"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	err := r.ParseForm()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	var pass []byte
	if !UserV.NameChecker(r.PostFormValue("user_name")) {
		ErrorHandler(w, r, http.StatusNotAcceptable)
		return
	}
	if !UserV.EmailChecker(r.PostFormValue("user_email")) {
		ErrorHandler(w, r, http.StatusNotAcceptable)
		return
	}
	// XXX What is best practice? validate frontend or backend?
	if !UserV.PasswordChecker(r.PostFormValue("password")) {
		ErrorHandler(w, r, http.StatusNotAcceptable)
		return
	} else {
		pass, err = bcrypt.GenerateFromPassword([]byte(r.PostFormValue("password")), 4)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	data := forum.Applicant{
		Username: r.PostFormValue("user_name"),
		Email:    r.PostFormValue("user_email"),
		Password: pass,
		Type:     "member",
	}
	if UserDb.UserDbRegisteration(data, db) != nil {
		ErrorHandler(w, r, http.StatusNotAcceptable)
		return
	}
	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "test.html", nil)
}
