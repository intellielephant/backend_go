/*
 * @Author: Will
 * @Date: 2022-11-14 18:53:38
 * @LastEditors: Will
 * @LastEditTime: 2022-11-14 19:11:41
 * @Description: 请填写简介
 */
package router

import (
	"backend-svc-go/controller"
	"backend-svc-go/global"
	"log"
	"runtime/debug"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	// router.Use(Recover)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	hello := controller.HelloWorld
	router.GET("/hello", hello)

	company := controller.NewCompanyController()
	router.POST("/littlein/register_company", company.RegisterCompany)
	router.POST("/littlein/register_store", company.RegisterStore)
	router.POST("/littlein/create_note", company.CreateNote)

	ai := controller.NewAIController()
	router.POST("/ai/send_message", ai.SendMessage)
	router.POST("/ai/get_function", ai.GetFunction)
	router.POST("/ai/quick_login", ai.QuickLogin)
	router.POST("/ai/get_baidu_access_token", ai.GetBaiduAccessToken)
	router.POST("/ai/user_phone_login", ai.UserPhoneLogin)

	router.POST("/ai/user_weixin_login", ai.UserWeixinLogin)

	// ai toolbox
	router.POST("/ai/get_category", ai.GetCategory)
	router.POST("/ai/get_hot", ai.GetHotApp)
	router.POST("/ai/get_app_by_categoryid", ai.GetAppByCategoryID)
	router.POST("/ai/update_avatar", ai.UpdateAvatar)
	router.POST("/ai/predict", ai.Predict)

	gemini := controller.NewGeminiController()
	router.POST("/gemini/gen_content", gemini.GenerateContent)
	router.POST("/gemini/summary_file", gemini.SummaryFile)

	router.POST("/gemini/uploadfile", gemini.UploadFile)
	router.POST("/gemini/list_file", gemini.ListFile)
	router.POST("/gemini/delete_file", gemini.DeleteFile)

	oss := controller.NewOSSController()
	router.POST("/oss/pushfile", oss.PushObject)

	fish := controller.NewFishController()
	router.POST("/fish/add_goods", fish.AddGoods)
	router.POST("/fish/get_goods", fish.GetGoods)
	router.POST("/fish/update_goods", fish.UpdateGoods)
	router.POST("/fish/add_order", fish.AddOrder)

	router.POST("/fish/get_category", fish.GetCategory)
	router.POST("/fish/get_order_by_table_name", fish.GetOrderByTableName)
	router.POST("/fish/get_orders", fish.GetOrders)

	return router
}

func HandleNotFound(ctx *gin.Context) {
	global.NewResult(ctx).Error(404, "资源未找到")
}

func Recover(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			global.NewResult(ctx).Error(500, "服务器内部错误")
		}
	}()
	ctx.Next()
}
