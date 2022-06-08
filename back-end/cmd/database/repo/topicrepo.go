package repo

import (
	"back-end/cmd/database/model"
	"context"
)

type CreateTopicIn struct {
	TopicName      string `json:"topic_name"`
	GroupStudentID string `json:"group_student_id"`
	LecturerID     string `json:"lecturer_id"`
	StartDay       string `json:"start_day"`
	EndDay         string `json:"end_day"`
	Allowance      string `json:"allowance"`
	TopicStatus    string `json:"topic_status"`
}

type CreateTopicOut struct {
	TopicID int64 `json:"topic_id"`
}

//Create Topic Repo for Topic repository

type TopicRepo interface {
	CreateTopic(context.Context, *CreateTopicIn) (*CreateTopicOut, error)
	FindByKeyWord(context.Context, string) (model.Topic, error)
	ListAll(ctx context.Context) ([]model.Topic, error)
}
