/*
 * @Author: Will
 * @Date: 2022-11-14 19:08:19
 * @LastEditors: Will
 * @LastEditTime: 2022-11-14 19:10:28
 * @Description: 请填写简介
 */
package controller

import (
	"backend-svc-go/global"

	"github.com/gin-gonic/gin"
)

func HelloWorld(ctx *gin.Context) {
	result := global.NewResult(ctx)
	result.Success("{'hello':'will'}")
}
