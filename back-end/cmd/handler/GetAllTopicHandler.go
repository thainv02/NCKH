package handler

import (
	"back-end/cmd/logic"
	"back-end/cmd/service"
	"back-end/core/http_response"
	"back-end/core/logger"
	"context"
	"github.com/gin-gonic/gin"
)

func GetAllTopicHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "trace_id", logger.GenerateTraceID("get-all-topic-api"))
		logHelper := logger.NewContextLog(ctx)
		getAllTopicLogic := logic.NewGetAllTopicLogic(ctx, logHelper, svcCtx)

		resp, err := getAllTopicLogic.GetAllTopic()
		if err != nil {
			panic(err)
		}

		http_response.ResponseJSON(c, resp.Code, resp)
	}
}