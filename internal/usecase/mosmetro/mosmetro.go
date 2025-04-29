package mosmetro

import (
	"mosmetro/internal/entity"
	"mosmetro/internal/repo"
)

type UseCase struct {
	stationRepo repo.StationRepo
}

func New(r repo.StationRepo) *UseCase {
	return &UseCase{stationRepo: r}
}

// add methods
func (us UseCase) GetStationById(id int) (entity.Station, error) {
	return us.stationRepo.GetStationById(id)
}

func (us UseCase) GetAllStations() ([]entity.Station, error) {
	return us.stationRepo.GetAllStations()
}
