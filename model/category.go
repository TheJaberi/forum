package forum

import (
	"context"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var CategoryError = errors.New("Error adding category!") //TODO Add remaining errors then move to errors file

func CreateCategory(name string) {
	err := createCategoryDb(name)
	if err != nil {
		log.Println(CategoryError.Error() + err.Error())
	}
	newCategory, err := GetCategory(name)
	if err != nil {
		log.Println(CategoryError.Error() + err.Error())
	}
	AllCategories = append(AllCategories, newCategory)
}

func createCategoryDb(name string) error {
	query := "INSERT INTO `Category` (`Name`) VALUES (?)"
	_, err := DB.ExecContext(context.Background(), query, name)
	if err != nil { // the category is added using the ExecContext
		return err
	}
	return nil
}

func GetCategories() {
	AllCategories = nil // FIXME Why do we clear it all the time?
	var category Category
	categoryData, _ := DB.Query("Select id, Name from Category") // FIXME Add proper error handling for query
	defer categoryData.Close()                                   // TODO when is defering required?
	for categoryData.Next() {
		categoryData.Scan(&category.CategoryID, &category.CategoryName) // FIXME Add proper error handling for scan
		AllCategories = append(AllCategories, category)
	}
}

func GetCategory(name string) (Category, error) {
	row := DB.QueryRow("SELECT id, Name from Category WHERE name=?", name) // FIXME Add proper error handling for query
	var category Category
	err := row.Scan(&category.CategoryID, &category.CategoryName) // FIXME Add proper error handling for scan
	if err != nil {
		log.Printf("Id not found")
		return category, err
	}
	return category, nil
}
