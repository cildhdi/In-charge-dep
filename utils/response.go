package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Ok = iota
	ParamError
	DatabaseError
	FailedAuthentication
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func response(ctx *gin.Context, code int, msg string, data interface{}) {
	if code == Ok {
		msg = "success"
	}
	ctx.AbortWithStatusJSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Success(ctx *gin.Context, data interface{}) {
	response(ctx, Ok, "", data)
}

func Error(ctx *gin.Context, code int, msg string) {
	response(ctx, code, msg, nil)
}
