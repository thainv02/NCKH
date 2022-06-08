package types

import "back-end/cmd/database/model"

type CreateTopicRequest struct {
	TopicName      string `json:"topic_name"`
	GroupStudentID string `json:"group_student_id"`
	LecturerID     string `json:"lecturer_id"`
	StartDay       string `json:"start_day"`
	EndDay         string `json:"end_day"`
	Allowance      string `json:"allowance"`
	TopicStatus    string `json:"topic_status"`
}

type CreateTopicResponse struct {
	Code  int          `json:"code"`
	Data  *model.Topic `json:"data"`
	Error string       `json:"error"`
}
