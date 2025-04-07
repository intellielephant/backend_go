/*
 * @Author: Will
 * @Date: 2022-11-14 18:46:55
 * @LastEditors: Will
 * @LastEditTime: 2022-11-14 18:51:39
 * @Description: 请填写简介
 */
package global

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Ctx *gin.Context
}

type ResultContent struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResult(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

func (r *Result) Success(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := ResultContent{}
	res.Code = 0
	res.Msg = ""
	res.Data = data
	r.Ctx.JSON(http.StatusOK, res)
}

func (r *Result) Error(code int, msg string) {
	res := ResultContent{}
	res.Code = code
	res.Msg = msg
	res.Data = gin.H{}
	r.Ctx.JSON(http.StatusOK, res)
}
