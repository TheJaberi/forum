package forum

import (
	"context"
	"log"
)

func CreateCategory(name string) {
	var cat Category
	query := "INSERT INTO `Category` (`Name`) VALUES (?)"
	_, err2 := DB.ExecContext(context.Background(), query, name)
	if err2 != nil { // the category is added using the ExecContext
		log.Fatal(err2)
	}
	cat.CategoryName = name
	AllCategories = append(AllCategories, cat)
}
