package controller

import (
	"backend-svc-go/global"
	"backend-svc-go/service"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

type AIController struct {
}

func NewAIController() AIController {
	return AIController{}
}

func (c *AIController) UserPhoneLogin(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("Parse form err ", err)
	}
	formData := make(map[string]interface{})
	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var phone string
	if value, ok := formData["phone"]; ok {
		phone = value.(string)
	}
	var invitor_code string
	if value, ok := formData["invitor_code"]; ok {
		invitor_code = value.(string)
	}
	user, err := service.UserPhoneLogin(phone, invitor_code)
	if err != nil {
		result.Error(400, "user phone login fail"+err.Error())
	} else {
		result.Success(gin.H{"user": user})
	}

}

func (c *AIController) UserWeixinLogin(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("Parse form err", err)
	}

	formData := make(map[string]interface{})
	json.NewDecoder(ctx.Request.Body).Decode(&formData)
	var code string

	if value, ok := formData["code"]; ok {
		code = value.(string)
	}

	var invitor_code string
	if value, ok := formData["invitor_code"]; ok {
		invitor_code = value.(string)
	}
	user, err := service.UserWeixinLogin(code, invitor_code)
	if err != nil {
		result.Error(400, "user weixin login fail"+err.Error())
	} else {
		result.Success(gin.H{"user": user})
	}

}

func (c *AIController) SendMessage(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var message string

	if value, ok := formData["message"]; ok {
		message = value.(string)
	}

	resp, err := service.SendMessage(message)
	if err != nil {
		result.Error(404, "send message fail"+err.Error())
	} else {
		result.Success(gin.H{"result": resp.Result})
	}
}

func (c *AIController) GetFunction(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var category string

	if value, ok := formData["category"]; ok {
		category = value.(string)
	}

	functions, err := service.GetFunction(category)
	if err != nil {
		result.Error(404, err.Error())
	} else {
		result.Success(gin.H{"functions": functions})
	}
}

// quick login
func (c *AIController) QuickLogin(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var token string

	if value, ok := formData["token"]; ok {
		token = value.(string)
	}

	var accessToken string

	if value, ok := formData["accessToken"]; ok {
		accessToken = value.(string)
	}

	isSuccess, info := service.OneClick(token, accessToken)
	if isSuccess {
		result.Success(gin.H{"phone": info})
	} else {
		result.Error(404, info)
	}
}

func (c *AIController) GetBaiduAccessToken(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}

	formData := make(map[string]interface{})
	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	access_token, err := service.GetBaiduAccessToken()
	if err != nil {
		result.Error(404, err.Error())
	} else {
		result.Success(gin.H{"access_token": access_token})
	}
}

func (c *AIController) GetCategory(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}

	formData := make(map[string]interface{})
	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	category, err := service.GetCategory()
	if err != nil {
		result.Error(404, err.Error())
	} else {
		result.Success(gin.H{"category": category})
	}
}

func (c *AIController) GetHotApp(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}

	formData := make(map[string]interface{})
	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	hotapp, err := service.GetHotApp()
	if err != nil {
		result.Error(404, err.Error())
	} else {
		result.Success(gin.H{"hotapp": hotapp})
	}
}

func (c *AIController) GetAppByCategoryID(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}

	formData := make(map[string]interface{})
	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var category_id int

	if value, ok := formData["category_id"]; ok {
		category_id = int(value.(float64))
	}

	app, err := service.GetAppByCategoryID(category_id)
	if err != nil {
		result.Error(404, err.Error())
	} else {
		result.Success(gin.H{"app": app})
	}
}

func (c *AIController) UpdateAvatar(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}

	formData := make(map[string]interface{})
	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var user_id int

	if value, ok := formData["user_id"]; ok {
		user_id = int(value.(float64))
	}

	var avatar string
	if value, ok := formData["avatar"]; ok {
		avatar = value.(string)
	}

	err = service.UpdateAvatar(user_id, avatar)
	if err != nil {
		result.Error(404, err.Error())
	} else {
		result.Success(gin.H{"success": true})
	}
}

func (c *AIController) Predict(ctx *gin.Context) {

	result := global.NewResult(ctx)

	err := ctx.Request.ParseForm()

	if err != nil {

		log.Println("parse form error ", err)

	}

	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	resp, err := service.Predict()

	if err != nil {

		result.Error(404, "send message fail"+err.Error())

	} else {

		result.Success(gin.H{"result": resp})

	}

}

func (c *AIController) ShenJia(ctx *gin.Context) {

	result := global.NewResult(ctx)

	err := ctx.Request.ParseForm()

	if err != nil {

		log.Println("parse form error ", err)

	}

	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	resp, err := service.ShenJia()

	if err != nil {

		result.Error(404, "send message fail"+err.Error())

	} else {

		result.Success(gin.H{"result": resp})

	}

}
