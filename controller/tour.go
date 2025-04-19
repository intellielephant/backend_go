package controller

import (
	"backend-svc-go/global"
	"backend-svc-go/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"backend-svc-go/service"

	"github.com/gin-gonic/gin"
)

type TourController struct{}

func NewTourController() TourController {
	return TourController{}
}

// CreateProduct handles the creation of a new product
func (pc *TourController) CreateProduct(ctx *gin.Context) {
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

	var description string
	if value, ok := formData["description"]; ok {
		description = value.(string)
	}

	var status int
	if value, ok := formData["status"]; ok {
		status = int(value.(float64))
	}

	var price float64
	if value, ok := formData["price"]; ok {
		price = value.(float64)
	}

	var category string
	if value, ok := formData["category"]; ok {
		category = value.(string)
	}

	var duration int
	if value, ok := formData["duration"]; ok {
		duration = int(value.(float64))
	}

	var image_url string
	if value, ok := formData["image_url"]; ok {
		image_url = value.(string)
	}

	layout := "2006-01-02"
	var departure_date string
	if value, ok := formData["departure_date"]; ok {
		departure_date = value.(string)
	}

	var return_date string
	if value, ok := formData["return_date"]; ok {
		return_date = value.(string)
	}

	var departure_location string
	if value, ok := formData["departure_location"]; ok {
		departure_location = value.(string)
	}

	var destination string
	if value, ok := formData["destination"]; ok {
		destination = value.(string)
	}

	var availability int
	if value, ok := formData["availability"]; ok {
		availability = int(value.(float64))
	}

	var product model.TourProducts
	product.Name = name
	product.Status = status
	product.Description = description
	product.Category = category
	product.Price = price
	product.Duration = duration
	product.ImageUrl = image_url
	product.Departure_date, _ = time.Parse(layout, departure_date)
	product.Return_date, _ = time.Parse(layout, return_date)
	product.Departure_location = departure_location
	product.Destination = destination
	product.Availability = availability
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	// 调用服务层保存产品
	if err := service.CreateProduct(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create product: " + err.Error(),
		})
		return
	}

	// 返回成功响应
	result.Success(gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}

// GetProductByID handles fetching a product by its ID
func (pc *TourController) GetProductByID(ctx *gin.Context) {
	result := global.NewResult(ctx)

	// 获取 URL 参数中的 id
	id := ctx.Param("id")

	// 将 id 转换为整数
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	// 查询产品
	product, err := service.GetProductByID(intID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found: " + err.Error(),
		})
		return
	}

	// 返回成功响应
	result.Success(gin.H{
		"message": "Product fetched successfully",
		"product": product,
	})
}

// GetProductsByCategory handles fetching products by category
func (pc *TourController) GetProductsByCategory(ctx *gin.Context) {
	result := global.NewResult(ctx)

	// 获取查询参数中的 category
	category := ctx.Query("category")
	if category == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Category parameter is required",
		})
		return
	}

	// 查询产品
	products, err := service.GetProductsByCategory(category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch products: " + err.Error(),
		})
		return
	}

	// 返回成功响应
	result.Success(gin.H{
		"message":  "Products fetched successfully",
		"products": products,
	})
}

// CreateCategory handles the creation of a new category
func (pc *TourController) CreateCategory(ctx *gin.Context) {
	result := global.NewResult(ctx)

	// 解析请求体
	var category model.TourCategory
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data: " + err.Error(),
		})
		return
	}

	// 设置创建时间和更新时间
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	// 调用服务层保存分类
	if err := service.CreateCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create category: " + err.Error(),
		})
		return
	}

	// 返回成功响应
	result.Success(gin.H{
		"message":  "Category created successfully",
		"category": category,
	})
}
