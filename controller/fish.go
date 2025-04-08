package controller

import (
	"backend-svc-go/global"
	"backend-svc-go/service"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type FishController struct {
}

func NewFishController() FishController {
	return FishController{}
}

func (c *FishController) AddGoods(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var name string

	if value, ok := formData["name"]; ok {
		name = value.(string)
	}

	var category int
	if value, ok := formData["category"]; ok {
		category = int(value.(float64))
	}

	var price float64
	if value, ok := formData["price"]; ok {
		price = value.(float64)
	}

	var image_url string
	if value, ok := formData["image_url"]; ok {
		image_url = value.(string)
	}

	goods, err := service.AddGoods(name, category, price, image_url)
	if err != nil {
		result.Error(404, "register company fail")
	} else {
		result.Success(gin.H{"goods": goods})
	}
}

func (c *FishController) GetGoods(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	goods_list, err := service.GetGoods()
	if err != nil {
		result.Error(404, "get goods fail")
	} else {
		result.Success(gin.H{"goods_list": goods_list})
	}
}

func (c *FishController) UpdateGoods(ctx *gin.Context) {
	result := global.NewResult(ctx)

	file, err := ctx.FormFile("file")

	if err != nil {
		result.Error(400, "get file fail"+err.Error())
	}

	src, err := file.Open()
	if err != nil {
		result.Error(400, "open file fail"+err.Error())
	}
	defer src.Close()
	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		result.Error(400, "read object fail"+err.Error())
		return
	}
	fileURL, err := service.PushObject(fileBytes, file.Filename)
	if err != nil {
		result.Error(400, "put object fail"+err.Error())
	} else {
		result.Success(gin.H{"file": fileURL})
	}
}

func (c *FishController) AddOrder(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var table_name string

	if value, ok := formData["table_name"]; ok {
		table_name = value.(string)
	}

	var goods_list []map[string]any
	if goodsListValue, ok := formData["goods_list"]; ok {
		if goodsList, ok := goodsListValue.([]interface{}); ok {
			for _, v := range goodsList {
				if goodsMap, ok := v.(map[string]interface{}); ok {
					goods_list = append(goods_list, goodsMap)
				}
			}
		}
	}

	order, order_goods, err := service.AddOrder(table_name, goods_list)
	if err != nil {
		result.Error(404, "add order fail")
	} else {
		result.Success(gin.H{"order": order, "order_goods": order_goods})
	}
}

func (c *FishController) GetOrderByTableName(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var table_name string

	if value, ok := formData["table_name"]; ok {
		table_name = value.(string)
	}

	order, order_goods, err := service.GetOrderByTableName(table_name)
	if err != nil {
		result.Error(404, "get order fail")
	} else {
		result.Success(gin.H{"order": order, "order_goods": order_goods})
	}
}

func (c *FishController) UpdateOrder(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var name string

	if value, ok := formData["name"]; ok {
		name = value.(string)
	}

	var category int
	if value, ok := formData["category"]; ok {
		category = value.(int)
	}

	var price float64
	if value, ok := formData["price"]; ok {
		price = value.(float64)
	}

	var image_url string
	if value, ok := formData["image_url"]; ok {
		image_url = value.(string)
	}

	goods, err := service.AddGoods(name, category, price, image_url)
	if err != nil {
		result.Error(404, "register company fail")
	} else {
		result.Success(gin.H{"goods": goods})
	}
}
