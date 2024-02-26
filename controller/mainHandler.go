package forum

import (
	model "forum/model"
	"html/template"
	"net/http"
)

// handles the main page
func MainHandler(w http.ResponseWriter, req *http.Request) {
	if model.AllData.IsLogged {
		check := model.CheckCookies(req)
		if check != nil {
			model.AllData.LoggedUser = model.Empty
			model.AllData.IsLogged = false
			model.LiveSession = model.EmptySession
		}
	}
	http.SetCookie(w, model.LoginCookie)
	if req.URL.Path != "/" {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		ErrorHandler(w, req, http.StatusMethodNotAllowed)
		return
	}
	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}
	model.GetCategories()
	model.GetPosts()
	model.AllData.AllPosts = model.RSort(model.AllPosts)
	model.AllData.AllCategories = model.AllCategories
	model.AllData.CategoryCheck = true
	model.SortPosts(req.FormValue("sortby"))
	model.FilterByCategory(req.FormValue("category"))
	if model.LoginError2 {
		model.AllData.LoginError = true
	} else {
		model.AllData.LoginError = false
	}
	if model.PostError2 {
		model.AllData.PostError = true
	} else {
		model.AllData.PostError = false
	}
	t.ExecuteTemplate(w, "index.html", model.AllData)
	model.AllData.LoginError = false
	model.LoginError2 = false
	model.AllData.PostError = false
	model.PostError2 = false
	model.AllData.LoginErrorMsg = ""
}
