package handler

import (
	"back-end/cmd/logic"
	"back-end/cmd/service"
	"back-end/cmd/types"
	"back-end/core/http_request"
	"back-end/core/http_response"
	"back-end/core/logger"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTopicHandler(svcCtx *service.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "trace_id", logger.GenerateTraceID("Create-topic-api"))
		logHelper := logger.NewContextLog(ctx)
		createTopicLogic := logic.NewCreateTopicLogic(ctx, svcCtx, logHelper)

		request := &types.CreateTopicRequest{}
		err := http_request.BindBodyJson(c, request)
		if err != nil {
			logHelper.Errorw("msg", "Failed while parse create topic request", "extra_readable_data", err.Error())
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		resp, err := createTopicLogic.CreateTopic(request)

		http_response.ResponseJSON(c, resp.Code, resp)
	}
}
