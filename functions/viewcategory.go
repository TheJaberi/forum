package forum

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func ViewCategory() {
	AllCategories = nil
	Database, errdatabase := sql.Open("sqlite3", "./forum.db")
	if errdatabase != nil {
		log.Fatal(errdatabase)
	}
	defer Database.Close()
	var category Category
	categoryData, _ := Database.Query("Select id, Name from Category")
	for categoryData.Next() {
		categoryData.Scan(&category.CategoryID, &category.CategoryName)
		defer categoryData.Close()
		AllCategories = append(AllCategories, category)
	}
}
