package model

import "gorm.io/gorm"

type AirlineInfo struct {
	gorm.Model
	Name string
	Icao *string
	Iata *string
	Logo *string
}

type AirlineInfoRepository struct {
	db *gorm.DB
}

func MakeAirlineInfoRepository(db *gorm.DB) AirlineInfoRepository {
	return AirlineInfoRepository{db}
}

func (m AirlineInfoRepository) Create(a AirlineInfo) {
	m.db.Create(&a)
}
