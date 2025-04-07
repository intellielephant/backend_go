package controller

import (
	"backend-svc-go/global"
	"backend-svc-go/service"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type OSSController struct {
}

func NewOSSController() OSSController {
	return OSSController{}
}

func (c *OSSController) PushObject(ctx *gin.Context) {
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
