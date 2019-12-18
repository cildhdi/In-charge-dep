package user

import (
	"github.com/gin-gonic/gin"

	"github.com/cildhdi/In-charge/models"
	"github.com/cildhdi/In-charge/utils"
)

func All(ctx *gin.Context) {
	var users []models.IcUser
	user := ctx.MustGet("user").(*models.IcUser)
	if user.Role == models.SuperUser {
		models.IcDb().Not("role", []int{models.SuperUser}).Find(&users)
	} else if user.Role == models.AdminUser {
		models.IcDb().Not("role", []int{models.AdminUser, models.SuperUser}).Find(&users)
	}
	utils.Success(ctx, users)
}
