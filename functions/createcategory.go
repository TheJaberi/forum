package forum

import (
	"context"
	"database/sql"
	"log"
)

func CreateCategory(name string) {
	var cat Category
	Database, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer Database.Close()
	query := "INSERT INTO `Category` (`Name`) VALUES (?)"
	_, err2 := Database.ExecContext(context.Background(), query, name)
	if err2 != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Fatal(err2)
	}
	cat.CategoryName = name
	AllCategories = append(AllCategories, cat)
}
