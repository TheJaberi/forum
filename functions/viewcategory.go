package forum

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)
func ViewCategory(){
	var count int
	AllCategories = nil
	Database, errdatabase := sql.Open("sqlite3", "./forum.db")
	if errdatabase != nil {
		log.Fatal(errdatabase)
	}
	defer Database.Close()
	countRows:= Database.QueryRow("SELECT COUNT(*) FROM Category")
	countRows.Scan(&count)
	fmt.Println(count)
	for i:=1;i<=count;i++{
	var category string
	categoryData := Database.QueryRow("Select Name from Category where id = ? AND post_id IS NULL", i)
	categoryData.Scan(&category)
	AllCategories = append(AllCategories, category)
	}
}
