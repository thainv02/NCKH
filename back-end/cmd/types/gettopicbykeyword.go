package types

import "back-end/cmd/database/model"

type GetTopicByKeyWordRequest struct {
	KeyWord string `json:"key_word"`
}

type GetTopicByKeyWordResponse struct {
	Code  int           `json:"code"`
	Data  []model.Topic `json:"data"`
	Error string        `json:"error"`
}
