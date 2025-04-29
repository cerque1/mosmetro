package httputils

import (
	"encoding/json"
)

type ErrorMessage struct {
	Code    int
	Message string
}

type BodyWithId struct {
	Id int `json:"id"`
}

func GetIdFromJson(body []byte) (BodyWithId, error) {
	var id BodyWithId
	err := json.Unmarshal(body, &id)
	if err != nil {
		return BodyWithId{}, err
	}
	return id, nil
}
