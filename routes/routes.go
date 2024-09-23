package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {

	basePath := "/api"

	docker := router.Group(basePath)
	{
		docker.GET("/containers", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Ok!",
				"code":    200,
			})
		})
	}
}
