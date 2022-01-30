package DAL

import (
	"cloudCharts/DAO"
	"cloudCharts/config"
)

func CreateDatabase(history DAO.History) DAO.History {
	tempHistory := DAO.History{
		Rank:         history.Rank,
		Name:         history.Name,
		Platform:     history.Platform,
		Year:         history.Year,
		Genre:        history.Genre,
		Publisher:    history.Publisher,
		NA_Sales:     history.NA_Sales,
		EU_Sales:     history.EU_Sales,
		JP_Sales:     history.JP_Sales,
		Other_Sales:  history.Other_Sales,
		Global_Sales: history.Global_Sales,
	}
	config.DB.Create(&tempHistory)
	return tempHistory
}

func GetDatabase() ([]DAO.History, error) {
	var histories []DAO.History
	err := config.DB.Model(&DAO.History{}).Find(&histories)
	if err.Error != nil {
		return nil, err.Error
	}
	return histories, nil
}
