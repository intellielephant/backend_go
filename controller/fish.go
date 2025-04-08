package controller

import (
	"backend-svc-go/global"
	"backend-svc-go/model"
	"backend-svc-go/service"
	"encoding/json"
	"fmt"
	"log"
	"time"

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
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var goods model.Goods
	if value, ok := formData["goods"]; ok {
		if goodsMap, ok := value.(map[string]interface{}); ok {
			goods, err = service.MapToGoods(goodsMap)
			if err != nil {
				result.Error(400, "invalid goods data")
				return
			}
		}
	}
	fmt.Println(goods)
	err = service.UpdateGoods(goods)
	if err != nil {
		result.Error(404, "register company fail"+err.Error())
	} else {
		result.Success(gin.H{"success": true})
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

func (c *FishController) GetOrders(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var start_date time.Time
	if value, ok := formData["start_date"]; ok {
		if dateStr, ok := value.(string); ok {
			layout := "2006-01-02 15:04:05" // 日期格式，例如 "YYYY-MM-DD"
			start_date, err = time.Parse(layout, dateStr)
			fmt.Println(start_date)
			if err != nil {
				result.Error(400, "Invalid date format")
				return
			}

		} else {
			fmt.Println(ok)
		}
	}

	var end_date time.Time

	if value, ok := formData["end_date"]; ok {
		if dateStr, ok := value.(string); ok {
			layout := "2006-01-02 15:04:05" // 日期格式，例如 "YYYY-MM-DD"
			end_date, err = time.Parse(layout, dateStr)
			fmt.Println(end_date)
			if err != nil {
				result.Error(400, "Invalid date format")
				return
			}

		}
	}
	fmt.Println(start_date, end_date)
	orders, err := service.GetOrders(start_date, end_date)
	if err != nil {
		result.Error(404, "get order fail")
	} else {
		result.Success(gin.H{"orders": orders})
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

	var id int
	if value, ok := formData["id"]; ok {
		id = int(value.(float64))
	}

	var status int
	if value, ok := formData["status"]; ok {
		status = int(value.(float64))
	}

	err = service.UpdateOrder(id, status)
	if err != nil {
		result.Error(404, "get order fail")
	} else {
		result.Success(gin.H{"success": true})
	}
}

func (c *FishController) GetCategory(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	categories, err := service.GetGoodsCategory()
	if err != nil {
		result.Error(404, "get order fail")
	} else {
		result.Success(gin.H{"categories": categories})
	}
}
