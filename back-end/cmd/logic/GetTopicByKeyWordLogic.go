package logic

import (
	"back-end/cmd/service"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type GetTopicByValue struct {
	ctx       context.Context
	svcCtx    *service.ServiceContext
	logHelper *log.Helper
}
