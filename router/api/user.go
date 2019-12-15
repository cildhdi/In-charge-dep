package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/cildhdi/In-charge/auth"
	"github.com/cildhdi/In-charge/models"
	"github.com/cildhdi/In-charge/utils"
)

type sendVerificationCodeBody struct {
	Phone string `json:"phone" binding:"required,len=11"`
}

func SendVerificationCode(ctx *gin.Context) {
	var param sendVerificationCodeBody
	if err := ctx.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		utils.Error(ctx, utils.ParamError, err.Error())
		return
	}

	vc := models.VerificationCode{
		Phone: param.Phone,
		Code:  1234,
	}

	if err := models.IcDb().Create(&vc).Error; err != nil {
		utils.Error(ctx, utils.DatabaseError, err.Error())
		return
	}

	utils.Success(ctx, nil)
}

var Login func(ctx *gin.Context)

func init() {
	Login = auth.GetMiddleware().LoginHandler
}
