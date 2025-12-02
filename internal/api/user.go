package api

import (
	storage "taskBoard_API/internal/repositories"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Repo *storage.UserRepository
}

func (h *UserHandler) Register(router *gin.Engine) {

}
