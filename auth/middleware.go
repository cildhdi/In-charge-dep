package auth

import (
	"errors"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/cildhdi/In-charge/config"
	"github.com/cildhdi/In-charge/models"
	"github.com/cildhdi/In-charge/utils"
)

const (
	identityKey = "phone"
	roleKey     = "role"
)

type loginBody struct {
	Phone string `json:"phone" binding:"required,len=11"`
	Code  string `json:"code" binding:"required,len=4"`
}

var middleware *jwt.GinJWTMiddleware

func init() {
	var err error
	middleware, err = jwt.New(
		&jwt.GinJWTMiddleware{
			Realm:       "ic",
			Key:         []byte(config.IcCfg().Auth.SecretKey),
			Timeout:     time.Hour * 24 * 7,
			IdentityKey: identityKey,
			PayloadFunc: func(data interface{}) jwt.MapClaims {
				if v, ok := data.(*models.IcUser); ok {
					return jwt.MapClaims{
						identityKey: v.Phone,
						roleKey:     v.Role,
					}
				} else {
					return jwt.MapClaims{}
				}
			},
			IdentityHandler: func(ctx *gin.Context) interface{} {
				claims := jwt.ExtractClaims(ctx)
				return &models.IcUser{
					Phone: claims[identityKey].(string),
					Role:  claims[roleKey].(*int),
				}
			},
			Authenticator: func(ctx *gin.Context) (interface{}, error) {
				var param loginBody
				if err := ctx.ShouldBindBodyWith(&param, binding.JSON); err != nil {
					return "", err
				}

				code, err := strconv.Atoi(param.Code)
				if err != nil || code < 1000 {
					return "", errors.New("code is not a number's string format or its value isnt in range")
				}

				var vc models.VerificationCode

				models.IcDb().Where(&models.VerificationCode{
					Phone: param.Phone,
					Code:  uint(code),
				}).First(&vc)
				if vc.ID == 0 {
					return nil, jwt.ErrFailedAuthentication
				}

				user := models.IcUser{
					Phone: vc.Phone,
				}
				models.IcDb().Where(&user).First(&user)

				if user.ID != 0 {
					return &user, nil
				} else {
					return nil, jwt.ErrFailedAuthentication
				}
			},
			Authorizator: func(data interface{}, ctx *gin.Context) bool {
				if v, ok := data.(*models.IcUser); ok && RoleCheck(ctx.Request.URL.Path, v.Role) {
					return true
				} else {
					return false
				}
			},
			Unauthorized: func(ctx *gin.Context, _ int, msg string) {
				utils.Error(ctx, utils.FailedAuthentication, msg)
			},
			TokenLookup:   "header: Authorization",
			TokenHeadName: "Bearer",
			TimeFunc:      time.Now,
		},
	)

	if err != nil {
		panic(err)
	}

}

func GetMiddleware() *jwt.GinJWTMiddleware {
	return middleware
}
