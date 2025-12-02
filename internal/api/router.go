package api

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(task *TaskHandler, user *UserHandler) *gin.Engine {

	r := gin.Default()

	task.Register(r)
	user.Register(r)
	return r

}
