package data

import (
	"encoding/json"

	"github.com/nsrvel/go-tools/models"
)

const (
	features = `
    [
        {
            "ID": 1,
            "Name": "Create new domain"
        },
        {
            "ID": 2,
            "Name": "Delete domain"
        },
        {
            "ID": 3,
            "Name": "Recover"
        }
    ]
    `
)

func ListFeature() (*[]models.Feature, error) {

	var listFeature []models.Feature

	err := json.Unmarshal([]byte(features), &listFeature)
	if err != nil {
		return nil, err
	}

	return &listFeature, nil
}
