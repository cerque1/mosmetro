package v1

import (
	"encoding/json"
	"io"
	"log"
	httputils "mosmetro/internal/http_utils"
	"net/http"
)

func (r *stationRoutes) getAllStations(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		if stations, err := r.s.GetAllStations(); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			val, _ := json.Marshal(httputils.ErrorMessage{Code: 500, Message: "Internal Server Error"})
			w.Write(val)
		} else {
			val, _ := json.Marshal(stations)
			w.Write(val)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		val, _ := json.Marshal(httputils.ErrorMessage{Code: 405, Message: "Method not allowed"})
		w.Write(val)
	}
}

func (r *stationRoutes) getStationById(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		body, err := io.ReadAll(req.Body)
		defer req.Body.Close()

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			val, _ := json.Marshal(httputils.ErrorMessage{Code: 500, Message: "Internal Server Error"})
			w.Write(val)
			return
		}

		id, err := httputils.GetIdFromJson(body)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			val, _ := json.Marshal(httputils.ErrorMessage{Code: 500, Message: "Internal Server Error"})
			w.Write(val)
			return
		}

		if station, err := r.s.GetStationById(id.Id); err != nil {
			w.WriteHeader(http.StatusNotFound)
			val, _ := json.Marshal(httputils.ErrorMessage{Code: 404, Message: "Not found"})
			w.Write(val)
		} else {
			val, _ := json.Marshal(station)
			w.Write(val)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		val, _ := json.Marshal(httputils.ErrorMessage{Code: 405, Message: "Method not allowed"})
		w.Write(val)
	}
}
