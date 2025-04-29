package v1

import (
	"fmt"
	"mosmetro/internal/usecase"
	"net/http"

	"github.com/go-chi/chi"
)

type stationRoutes struct {
	s usecase.Stations
}

func NewRouter(us usecase.Stations) http.Handler {
	routes := stationRoutes{s: us}
	router := chi.NewRouter()

	router.HandleFunc("/Hello", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hello") })
	router.HandleFunc("/get_all_stations", routes.getAllStations)
	router.HandleFunc("/find_station_by_id", routes.getStationById)

	return router
}
