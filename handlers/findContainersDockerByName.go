package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *DockerHandler) FindContainersByImage(ctx *gin.Context) {

	imageName := ctx.Query("image_name")

	if imageName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":        "O nome da imagem é obrigatório",
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	containers, err := h.UseCase.FindContainersByImage(imageName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg":        "Não foi encontrado nenhum container com esse nome de Imagem",
			"statusCode": http.StatusNotFound,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"containers": containers,
	})
}
