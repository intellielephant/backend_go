package service

import (
	"backend-svc-go/dao"
	"backend-svc-go/model"
)

// CreateProduct 保存产品到数据库
func CreateProduct(product *model.TourProducts) error {
	// 使用 GORM 的 Create 方法将产品插入数据库
	return dao.CreateProduct(product)
}

func GetProductByID(id int) (*model.TourProducts, error) {
	product, err := dao.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
