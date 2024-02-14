package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
)

func create_copies(amount int, book_id int) error {
	for i := 0; i < amount; i++ {
		var copy models.Copy
		copy.BookID = uint(book_id)
		copy.IsBorrowed = false

		err := database.DB.Create(copy).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func delete_copies(amount int, book_id int) error {
	var deletable_amount int64
	database.DB.Model(&models.Copy{}).Where("book_id = ? AND is_borrowed = ?", book_id, false).Count(&deletable_amount)

	if amount > int(deletable_amount) {
		amount = int(deletable_amount)
	}

	for i := 0; i < amount; i++ {
		var copy_to_delete models.Copy

		err := database.DB.Where("book_id = ? AND is_borrowed = ?", book_id, false).First(&copy_to_delete).Error
		if err != nil {
			return err
		}

		err = database.DB.Delete(&copy_to_delete).Error
		if err != nil {
			return err
		}
	}

	return nil
}
