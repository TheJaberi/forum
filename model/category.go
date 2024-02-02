package forum

import (
	"context"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	NewCategoryError  = errors.New("Error adding category!")
	ScanCategoryError = errors.New("Categry Scan Error!")
)

func CreateCategory(name string) {
	err := createCategoryDb(name)
	if err != nil {
		log.Println(NewCategoryError.Error() + err.Error())
	}
	newCategory, err := GetCategory(name)
	if err != nil {
		log.Println(NewCategoryError.Error() + err.Error())
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

func GetCategories() error {
	AllCategories = nil // FIXME Why do we clear it all the time?
	var category Category
	categoryData, err := DB.Query("Select id, Name from Category")
	if err != nil {
		return err
	}
	defer categoryData.Close()
	for categoryData.Next() {
		err := categoryData.Scan(&category.CategoryID, &category.CategoryName)
		if err != nil {
			log.Printf(ScanCategoryError.Error())
			return err
		}
		AllCategories = append(AllCategories, category)
	}
	return nil
}

func GetCategory(name string) (Category, error) {
	row := DB.QueryRow("SELECT id, Name from Category WHERE name=?", name)
	var category Category
	err := row.Scan(&category.CategoryID, &category.CategoryName)
	if err != nil {
		log.Printf(ScanCategoryError.Error())
		return category, err
	}
	return category, nil
}
