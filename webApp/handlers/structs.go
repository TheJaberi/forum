package forum

type Login struct {
	User_email string
	User_pass  string
}

type Register struct {
	User_name  string
	User_email string
	User_pass  string
	User_type  string
}
