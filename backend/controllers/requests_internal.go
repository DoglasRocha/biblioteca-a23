package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
)

func get_open_requests() ([]models.Request, error) {
	var requests []models.Request

	err := database.DB.Where("is_accepted = ?", false).Find(&requests).Error
	if err != nil {
		return []models.Request{}, err
	}

	for i := range requests {
		err = database.PopulateRequest(&requests[i], requests[i].ID)
		if err != nil {
			return []models.Request{}, err
		}
	}

	return requests, nil
}
