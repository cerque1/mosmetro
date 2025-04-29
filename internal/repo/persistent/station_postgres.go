package persistent

import (
	"database/sql"
	"log"
	"mosmetro/internal/entity"
)

type StationRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *StationRepo {
	return &StationRepo{db: db}
}

func (sr StationRepo) GetStationById(id int) (entity.Station, error) {
	row := sr.db.QueryRow("select * from stations where id = $1", id)
	station := entity.Station{}
	err := row.Scan(&station.Id, &station.Name, &station.Branch)
	if err != nil {
		return entity.Station{}, err
	}
	return station, nil
}

func (sr StationRepo) GetAllStations() ([]entity.Station, error) {
	rows, err := sr.db.Query("select * from stations")
	if err != nil {
		return []entity.Station{}, err
	}
	defer rows.Close()
	stations := []entity.Station{}

	for rows.Next() {
		s := entity.Station{}
		err := rows.Scan(&s.Id, &s.Name, &s.Branch)
		if err != nil {
			log.Fatalln(err)
			continue
		}
		stations = append(stations, s)
	}

	return stations, nil
}
