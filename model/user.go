package forum

import (
	"database/sql"
)

// Insert new row into user table
func userInsertDb(applicant Applicant, db *sql.DB, pass []byte) error {
	sqlStmt, err := db.Prepare("INSERT INTO users (user_name, user_email, user_pass, user_type) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = sqlStmt.Exec(applicant.Username, applicant.Email, pass, applicant.Type)
	if err != nil {
		return err
	}
	return nil
}

// Check if email exists in user table
func UserExistsDb(applicantEmail string) error {
	sqlStmt := `SELECT EXISTS (SELECT 1 FROM users WHERE user_email = ?)`
	var exists bool
	err := DB.QueryRow(sqlStmt, applicantEmail).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return UserExistsError
	}
	return nil
}

// Retrieve data from the user table and assign into the global struct
func UserRetrieveDb(email string, password string) error {
	userdata := DB.QueryRow("SELECT user_id, user_name, user_pass, user_email, user_type FROM users where user_email = ?", email) // select gets the data from users table
	err := userdata.Scan(&LoggedUser.Userid, &LoggedUser.Username, &LoggedUser.Password, &LoggedUser.Email, &LoggedUser.Type)     // scan assigns the data of the row to variables
	if err != nil {
		return err
	}
	return nil
}
