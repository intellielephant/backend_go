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
	"encoding/json"
	"log"

	"backend-svc-go/service"

	"github.com/gin-gonic/gin"
)

type CompanyController struct{}

func NewCompanyController() CompanyController {
	return CompanyController{}
}

func (c *CompanyController) RegisterCompany(ctx *gin.Context) {
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

	var contact_name string
	if value, ok := formData["contact_name"]; ok {
		contact_name = value.(string)
	}

	var phone string
	if value, ok := formData["phone"]; ok {
		phone = value.(string)
	}

	var address string
	if value, ok := formData["address"]; ok {
		address = value.(string)
	}

	company, err := service.RegisterCompany(name, contact_name, phone, address)
	if err != nil {
		result.Error(404, "register company fail")
	} else {
		result.Success(gin.H{"company": company})
	}
}

func (c *CompanyController) RegisterStore(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var company_no string

	if value, ok := formData["company_no"]; ok {
		company_no = value.(string)
	}

	var name string

	if value, ok := formData["name"]; ok {
		name = value.(string)
	}

	var contact_name string
	if value, ok := formData["contact_name"]; ok {
		contact_name = value.(string)
	}

	var phone string
	if value, ok := formData["phone"]; ok {
		phone = value.(string)
	}

	var address string
	if value, ok := formData["address"]; ok {
		address = value.(string)
	}

	store, err := service.RegisterStore(company_no, name, contact_name, phone, address)
	if err != nil {
		result.Error(404, "register store fail")
	} else {
		result.Success(gin.H{"store": store})
	}
}

func (c *CompanyController) CreateNote(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var store_id int

	if value, ok := formData["store_id"]; ok {
		store_id = int(value.(float64))
	}

	var name string

	if value, ok := formData["name"]; ok {
		name = value.(string)
	}

	note, err := service.CreateNote(store_id, name)
	if err != nil {
		result.Error(404, "create note fail")
	} else {
		result.Success(gin.H{"note": note})
	}
}

func (c *CompanyController) CreateGoods(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var store_id int

	if value, ok := formData["store_id"]; ok {
		store_id = int(value.(float64))
	}

	var note_id int

	if value, ok := formData["note_id"]; ok {
		note_id = int(value.(float64))
	}

	var title string

	if value, ok := formData["title"]; ok {
		title = value.(string)
	}

	goods, err := service.CreateGoods(store_id, note_id, title)
	if err != nil {
		result.Error(404, "create goods fail")
	} else {
		result.Success(gin.H{"goods": goods})
	}
}

func (c *CompanyController) SelectGoods(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var page_size int

	if value, ok := formData["page_size"]; ok {
		page_size = int(value.(float64))
	}

	var page_index int

	if value, ok := formData["page_index"]; ok {
		page_index = int(value.(float64))
	}

	goods, err := service.SelectGoods(page_size, page_index)
	if err != nil {
		result.Error(404, "create goods fail")
	} else {
		result.Success(goods)
	}
}
