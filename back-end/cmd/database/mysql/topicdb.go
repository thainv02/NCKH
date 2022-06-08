package mysql

import (
	"back-end/cmd/database/model"
	"back-end/cmd/database/repo"
	"back-end/core/logger"
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TopicDB struct {
	table   string
	connect *sqlx.DB
}

func NewTopicDB() (*TopicDB, error) {
	db, err := sqlx.Open("mysql", "root:524020@tcp(127.0.0.1:3306)/nckh")
	if err != nil {
		panic(err)
	}
	return &TopicDB{
		table:   "topic",
		connect: db,
	}, nil
}

func (topic *TopicDB) Close() {
	err := topic.connect.Close()
	if err != nil {
		return
	}
}

func (topic *TopicDB) CreateTopic(ctx context.Context, input *repo.CreateTopicIn) (*repo.CreateTopicOut, error) {
	ctxLogger := logger.NewContextLog(ctx)
	// ctxLogger to write err
	db := sq.Insert(topic.table).Columns(GetListColumn(input)...).Values(GetListValues(input)...)
	query, arg, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query - err: %s", err.Error())
		return nil, err
	}
	result := topic.connect.MustExec(query, arg...)
	insertedID, err := result.LastInsertId()
	if err != nil {
		ctxLogger.Errorf("Failed while insert topic to database, err: %s", err.Error())
		return nil, err
	}
	return &repo.CreateTopicOut{
		TopicID: insertedID,
	}, nil
}

func (topic *TopicDB) FindByKeyWord(ctx context.Context, keyWord string) (model.Topic, error) {
	ctxLogger := logger.NewContextLog(ctx)
	var result []model.Topic

	db := sq.Select("*").From(topic.table)
	query, arg, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query, err: %s", err.Error())
		return result[0], err
	}
	err = topic.connect.Select(&result, query, arg...)
	if err != nil {
		if err == sql.ErrNoRows {
			ctxLogger.Error("Don't have any topic")
			return result[0], err
		}
		ctxLogger.Errorf("Failed while select topic, err: %s", err.Error())
		return result[0], err
	}
	return result[0], nil
}

func (topic *TopicDB) ListAll(ctx context.Context) ([]model.Topic, error) {
	ctxLogger := logger.NewContextLog(ctx)
	var result []model.Topic

	db := sq.Select("*").From(topic.table)
	query, arg, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query, err: %s", err.Error())
		return nil, err
	}
	err = topic.connect.Select(&result, query, arg...)
	if err != nil {
		if err == sql.ErrNoRows {
			ctxLogger.Error("Don't have any topic")
			return nil, err
		}
		ctxLogger.Errorf("Failed while select topic, err: %s", err.Error())
		return nil, err
	}
	return result, nil
}
