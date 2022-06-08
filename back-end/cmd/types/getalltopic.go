package types

import "back-end/cmd/database/model"

type GetAllTopicResponse struct {
	Code  int           `json:"code"`
	Data  []model.Topic `json:"data"`
	Error string        `json:"error"`
}
