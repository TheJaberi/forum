package forum

import (
	_ "github.com/mattn/go-sqlite3"
)

func ViewCategory() {
	AllCategories = nil
	var category Category
	categoryData, _ := DB.Query("Select id, Name from Category")
	for categoryData.Next() {
		categoryData.Scan(&category.CategoryID, &category.CategoryName)
		defer categoryData.Close()
		AllCategories = append(AllCategories, category)
	}
}
