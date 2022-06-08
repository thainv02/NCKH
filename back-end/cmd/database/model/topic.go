package model

type Topic struct {
	TopicID        string `db:"topic_id"`
	TopicName      string `db:"topic_name"`
	GroupStudentID string `db:"group_student_id"`
	LecturerID     string `db:"lecturer_id"`
	StartDay       string `db:"start_day"`
	EndDay         string `db:"end_day"`
	Allowance      string `db:"allowance"`
	TopicStatus    string `db:"topic_status"`
}
