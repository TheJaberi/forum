package forum

import (
	"context"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// TABLE: Category

func CreateCategory(name string) {
	err := CreateCategoryDb(name)
	if err != nil {
		log.Println(NewCategoryError.Error() + err.Error())
	}
	newCategory, err := GetCategory(name)
	if err != nil {
		log.Println(NewCategoryError.Error() + err.Error())
	}
	AllCategories = append(AllCategories, newCategory)
}

func CreateCategoryDb(name string) error {
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

// TABLE: Post2Category

func AssignPostCategoryDb(postID int64, postCategories []int) error {
	for _, category := range postCategories {
		queryCategory := "INSERT INTO `Post2Category` (`post_id`, `category_id`) VALUES (?, ?)"
		_, err := DB.ExecContext(context.Background(), queryCategory, postID, category)
		if err != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
			log.Fatal(err)
			return err
		}
		/*
			for j := 0; j < len(AllCategories); j++ {
				if AllCategories[j].CategoryID == postCategories[i] {
					postData.Category = append(postData.Category, AllCategories[j])
					break
				}
			}
		*/
	}
	err := GetCategories()
	if err != nil {
		return err
	}
	return nil
}

func GetPostCategories(p *Post) error {
	categoryData, err := DB.Query("Select category_id from Post2Category where post_id = ?", p.PostID)
	if err != nil {
		return errors.New("Category Query Error:" + err.Error())
	}
	defer categoryData.Close()
	for categoryData.Next() {
		var categoryID int
		err := categoryData.Scan(&categoryID)
		if err != nil {
			return errors.New("Category Scan Error:" + err.Error())
		}
		for i := range AllCategories {
			if categoryID == AllCategories[i].CategoryID {
				p.Category = append(p.Category, AllCategories[i])
				break
			}
		}
	}
	return nil
}
