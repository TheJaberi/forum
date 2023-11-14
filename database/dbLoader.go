package forum

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func DatabaseLoader() {
	var err error
	DB, err = sql.Open("sqlite3", "../database/forum.db")
	if err != nil {
		log.Fatalf("%v", err)
	}
	if !IsTableExists(DB, "users") {
		DataBase(DB)
		sqlStmt, err := DB.Prepare("INSERT INTO users (user_name, user_email, user_pass, user_type) VALUES (?, ?, ?, ?)")
		if err != nil {
			log.Fatalf("%v", err)
		}
		sqlStmt.Exec("admin", "admin", "admin", "admin")
	}

}

func IsTableExists(db *sql.DB, tableName string) bool {
	sqlStmt, err := db.Prepare("SELECT COUNT(*) FROM sqlite_master WHERE type = 'table' AND name = ?")
	if err != nil {
		return false
	}
	defer sqlStmt.Close()
	var count int
	err = sqlStmt.QueryRow(tableName).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}
