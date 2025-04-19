/*
 * @Author: Will
 * @Date: 2022-11-10 11:15:12
 * @LastEditors: Will
 * @LastEditTime: 2022-11-14 19:07:21
 * @Description: 请填写简介
 */
package main

import (
	"backend-svc-go/global"
	"backend-svc-go/router"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {
	err := global.SetupDBAI()
	if err != nil {
		log.Fatalf("init setup db engine err: %v", err)
		return
	}

	err = global.SetupDBFish()
	if err != nil {
		log.Fatalf("init setup db engine err: %v", err)
		return
	}

	err = global.SetupDBTour()
	if err != nil {
		log.Fatalf("init setup db engine err: %v", err)
		return
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "littleadds" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization"})
			ctx.Abort()
			return
		}
	}
}

func main() {
	Init()
	r := router.Router()
	r.Use(AuthMiddleware())
	r.Run(":8080")
}
