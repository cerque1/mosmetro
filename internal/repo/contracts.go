package repo

import "mosmetro/internal/entity"

type StationRepo interface {
	GetStationById(id int) (entity.Station, error)
	GetAllStations() ([]entity.Station, error)
}
