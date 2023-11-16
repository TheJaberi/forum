package forum

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func ViewCategory() {
	AllCategories = nil
	Database, errdatabase := sql.Open("sqlite3", "./forum.db")
	if errdatabase != nil {
		log.Fatal(errdatabase)
	}
	defer Database.Close()
		var category Category
		categoryData, _ := Database.Query("Select id, Name from Category where post_id IS NULL")
		for categoryData.Next(){
		categoryData.Scan(&category.CategoryID, &category.CategoryName)
		fmt.Println(category)
		defer categoryData.Close()
		AllCategories = append(AllCategories, category)}
	}
