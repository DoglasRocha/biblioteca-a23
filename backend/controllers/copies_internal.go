package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"log/slog"
)

func create_copies(amount int, book_id int) error {
	for i := 0; i < amount; i++ {
		err := database.DB.Create(&models.Copy{BookID: uint(book_id), IsBorrowed: false}).Error
		if err != nil {
			slog.Warn(
				"Erro ao criar cópia",
				"err", err,
			)
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
			slog.Warn(
				"Erro ao encontrar cópia disponível para deleção",
				"err", err,
			)
			return err
		}

		err = database.DB.Delete(&copy_to_delete).Error
		if err != nil {
			slog.Warn(
				"Erro ao deletar cópia",
				"err", err,
				"copy_id", copy_to_delete.ID,
			)
			return err
		}
	}

	return nil
}
