package forum
import "database/sql"
var (
	Database *sql.DB
	LoggedUser User
	ErrorMsg string
)
type User struct{
	Userid int
	Username string
	Password string
	Email string
	Registered bool
}
var ErrResponse struct {
	StatusCode bool
	ErrorMsg string
}
