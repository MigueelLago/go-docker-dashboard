package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *DockerHandler) StopContainerDocker(ctx *gin.Context) {

	containeID := ctx.Param("containerID")
	if containeID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":        "O ID do container é necessáiro",
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	message, err := h.UseCase.StopContainerDocker(containeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":        "Houve um erro ao tentar parar o container",
			"statusCode": http.StatusInternalServerError,
			"error":      err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":        message,
		"statusCode": 200,
	})
}
