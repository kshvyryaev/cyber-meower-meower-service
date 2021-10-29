package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/command"
)

type MeowController struct {
	CreateMeowCommandHandler *command.СreateMeowCommandHandler
}

func (controller *MeowController) Route(router *gin.Engine) {
	meow := router.Group("/meow")
	{
		meow.POST("/", controller.CreateMeow)
	}
}

func (controller *MeowController) CreateMeow(context *gin.Context) {
	var command command.CreateMeowCommand
	if err := context.BindJSON(&command); err != nil {
		// TODO: Добавить обработку ошибок
		return
	}

	response, err := controller.CreateMeowCommandHandler.Handle(&command)
	if err != nil {
		// TODO: Добавить обработку ошибок
		return
	}

	context.JSON(http.StatusOK, response)
}

func ProvideMeowController(createMeowCommandHandler *command.СreateMeowCommandHandler) *MeowController {
	return &MeowController{
		CreateMeowCommandHandler: createMeowCommandHandler,
	}
}

var MeowControllerSet = wire.NewSet(ProvideMeowController)
