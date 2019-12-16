package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"

	"github.com/cildhdi/In-charge/auth"
	"github.com/cildhdi/In-charge/models"
	"github.com/cildhdi/In-charge/utils"
)

func Status(ctx *gin.Context) {
	utils.Success(ctx, nil)
}

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

type registerBody struct {
	Phone string `json:"phone" binding:"required,len=11"`
	Role  string `json:"code" binding:"required,gte=0"`
}

func SuperRegister(ctx *gin.Context) {
	var param registerBody
	if err := ctx.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		utils.Error(ctx, utils.ParamError, err.Error())
		return
	}

	var user models.IcUser
	if v, ok := ctx.Get("user"); ok {
		if userValue, ok := v.(*models.IcUser); ok {
			user = *userValue
		}
	} else {
		utils.Error(ctx, utils.FailedAuthentication, "failed to identify your account")
		return
	}

	if !(*user.Role == models.SuperUser || *user.Role == models.AdminUser) {
		utils.Error(ctx, utils.FailedAuthentication, "you don't have permission to access this resource")
		return
	}

	if role, err := strconv.Atoi(param.Role); err == nil {
		newUser := models.IcUser{
			Phone: param.Phone,
			Role:  &role,
		}
		if err := models.IcDb().Create(&newUser).Error; err != nil {
			utils.Error(ctx, utils.DatabaseError, err.Error())
			return
		} else {
			utils.Success(ctx, param)
			return
		}
	} else {
		utils.Error(ctx, utils.ParamError, err.Error())
		return
	}
}

func init() {
	Login = auth.GetMiddleware().LoginHandler
}
