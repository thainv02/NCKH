package logic

import (
	"back-end/cmd/service"
	"back-end/cmd/types"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
)

type GetAllTopicLogic struct {
	ctx       context.Context
	svcCtx    *service.ServiceContext
	logHelper *log.Helper
}

func NewGetAllTopicLogic(ctx context.Context, logHelper *log.Helper, svcCtx *service.ServiceContext) GetAllTopicLogic {
	return GetAllTopicLogic{
		ctx:       ctx,
		logHelper: logHelper,
		svcCtx:    svcCtx,
	}
}

func (l *GetAllTopicLogic) GetAllTopic() (*types.GetAllTopicResponse, error) {
	//l.logHelper.Info("Start greet: ")

	topics, err := l.svcCtx.TopicRepo.ListAll(l.ctx)
	if err != nil {
		l.logHelper.Errorf("Failed while get topic, err: %s", err.Error())
		return &types.GetAllTopicResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		}, err
	}
	return &types.GetAllTopicResponse{
		Code: http.StatusOK,
		Data: topics,
	}, nil
}
