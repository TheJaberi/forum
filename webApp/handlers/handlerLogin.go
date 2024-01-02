package forum

import (
	"database/sql"
	"fmt"
	"forum"
	UserDb "forum/database"
	v "forum/webApp/validators"
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	err := r.ParseForm()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	for _, yup := range r.Form {
		fmt.Println(yup)
	}
	for _, aha := range r.Header {
		fmt.Println(aha)
	}

	data := forum.Login{
		Email:    r.PostFormValue("user_email2"),
		Password: r.PostFormValue("password2"),
	}
	if UserDb.UserDbLogin(data, db) != nil {
		ErrorHandler(w, r, http.StatusNotAcceptable)
		return
	}

	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:     "exampleCookie",
		Value:    "Hello world!",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	err = v.WriteSignedCookie(w, cookie, v.SecretKey)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "test.html", nil)
}
