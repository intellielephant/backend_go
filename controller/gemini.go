package controller

import (
	"backend-svc-go/global"
	"backend-svc-go/service"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type GeminiController struct{}

func NewGeminiController() GeminiController {
	return GeminiController{}
}

func (c *GeminiController) GenerateContent(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)
	var prompt string

	if value, ok := formData["prompt"]; ok {
		prompt = value.(string)
	}

	var tool string

	if value, ok := formData["tool"]; ok {
		tool = value.(string)
	}

	resp, err := service.GenerateContent2(prompt, tool)
	if err != nil {
		result.Error(400, err.Error())
	} else {
		result.Success(gin.H{"resp": resp})
	}

}

func (c *GeminiController) SummaryFile(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)
	var filename string

	if value, ok := formData["filename"]; ok {
		filename = value.(string)
	}

	var prompt string

	if value, ok := formData["prompt"]; ok {
		prompt = value.(string)
	}

	resp, err := service.DescribeFile(filename, prompt)
	if err != nil {
		result.Error(400, err.Error())
	} else {
		result.Success(gin.H{"resp": resp})
	}
}

func (c *GeminiController) ListFile(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	files, err := service.ListFile()
	if err != nil {
		result.Error(400, err.Error())
	} else {
		result.Success(gin.H{"files": files})
	}
}

func (c *GeminiController) DeleteFile(ctx *gin.Context) {
	result := global.NewResult(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Println("parse form error ", err)
	}
	formData := make(map[string]interface{})

	json.NewDecoder(ctx.Request.Body).Decode(&formData)

	var filename string

	if value, ok := formData["filename"]; ok {
		filename = value.(string)
	}

	err = service.DeleteFile(filename)
	if err != nil {
		result.Error(400, err.Error())
	} else {
		result.Success(gin.H{"success": true})
	}
}

func (c *GeminiController) UploadFile(ctx *gin.Context) {
	result := global.NewResult(ctx)
	file, err := ctx.FormFile("file")
	if err != nil {
		result.Error(400, err.Error())
		return
	}

	filePath := filepath.Join(os.Getenv("uploadsPath"), file.Filename)

	err = ctx.SaveUploadedFile(file, filePath)

	if err != nil {
		result.Error(400, err.Error())
		return
	}
	print(file.Header)
	genFile, err := service.UploadFile(filePath, file.Filename)

	if err != nil {
		result.Error(400, err.Error())
	} else {
		result.Success(gin.H{"file": genFile})
	}
}
