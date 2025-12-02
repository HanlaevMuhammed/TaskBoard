package api

import (
	storage "taskBoard_API/internal/repositories"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	Repo *storage.TaskRepository
}

func (h *TaskHandler) Register(router *gin.Engine) {

}
