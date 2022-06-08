package logic

import (
	"back-end/cmd/database/model"
	"back-end/cmd/database/repo"
	"back-end/cmd/service"
	"back-end/cmd/types"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
	"strconv"
)

type CreateTopicLogic struct {
	ctx       context.Context
	svcCtx    *service.ServiceContext
	logHelper *log.Helper
}

func NewCreateTopicLogic(ctx context.Context, svcCtx *service.ServiceContext, logHelper *log.Helper) CreateTopicLogic {
	return CreateTopicLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *CreateTopicLogic) CreateTopic(input *types.CreateTopicRequest) (*types.CreateTopicResponse, error) {
	//l.logHelper.Infof("Start process greet message, input: %+v", input)
	data := &repo.CreateTopicIn{
		TopicName:      input.TopicName,
		GroupStudentID: input.GroupStudentID,
		LecturerID:     input.LecturerID,
		StartDay:       input.StartDay,
		EndDay:         input.EndDay,
		Allowance:      input.Allowance,
		TopicStatus:    input.TopicStatus,
	}

	resp, err := l.svcCtx.TopicRepo.CreateTopic(l.ctx, data)
	if err != nil {
		return &types.CreateTopicResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		}, err
	}

	return &types.CreateTopicResponse{
		Code: http.StatusOK,
		Data: &model.Topic{
			TopicID:        strconv.FormatInt(resp.TopicID, 10),
			TopicName:      data.TopicName,
			GroupStudentID: data.GroupStudentID,
			LecturerID:     data.LecturerID,
			StartDay:       data.StartDay,
			EndDay:         data.EndDay,
			Allowance:      data.Allowance,
			TopicStatus:    data.TopicStatus,
		},
	}, nil
}
