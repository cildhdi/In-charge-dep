package router

import (
	"github.com/gin-gonic/gin"

	"github.com/cildhdi/In-charge/utils"
)

func Status(ctx *gin.Context) {
	utils.Success(ctx, nil)
}
