package usecase

import "mosmetro/internal/entity"

type Stations interface {
	GetStationById(id int) (entity.Station, error)
	GetAllStations() ([]entity.Station, error)
}
