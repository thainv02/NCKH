package handler

import (
	"back-end/cmd/service"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(router *gin.Engine, serverCtx *service.ServiceContext) {
	router.GET("/get_all_topic", GetAllTopicHandler(serverCtx))
	router.POST("/create_topic", CreateTopicHandler(serverCtx))
}
