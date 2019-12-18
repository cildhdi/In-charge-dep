package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

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

type superRegisterBody struct {
	Phone string `json:"phone" binding:"required,len=11"`
	Role  int    `json:"role" binding:"min=0,max=2"`
	Name  string `json:"name" binding:"required"`
}

func SuperRegister(ctx *gin.Context) {
	var param superRegisterBody
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

	if !(user.Role == models.SuperUser || user.Role == models.AdminUser) {
		utils.Error(ctx, utils.FailedAuthentication, "you don't have permission to access this resource")
		return
	}

	newUser := models.IcUser{
		Phone: param.Phone,
		Role:  param.Role,
		Name:  param.Name,
	}
	if err := models.IcDb().Create(&newUser).Error; err != nil {
		utils.Error(ctx, utils.DatabaseError, err.Error())
		return
	} else {
		utils.Success(ctx, param)
		return
	}
}

func Register(ctx *gin.Context) {
	var param auth.PhoneCodeBody
	if err := ctx.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		utils.Error(ctx, utils.ParamError, err.Error())
		return
	}

	if err := utils.CodeVerify(param.Phone, param.Code, true); err != nil {
		utils.Error(ctx, utils.FailedCodeVerify, err.Error())
		return
	}

	user := models.IcUser{
		Phone: param.Phone,
		Role:  models.CustomerUser,
	}

	if err := models.IcDb().Create(&user).Error; err != nil {
		utils.Error(ctx, utils.DatabaseError, err.Error())
	} else {
		utils.Success(ctx, &user)
	}
}

func init() {
	Login = auth.GetMiddleware().LoginHandler
}
