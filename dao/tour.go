package dao

import (
	"backend-svc-go/global"
	"backend-svc-go/model"
)

// CreateProduct 将产品保存到数据库
func CreateProduct(product *model.TourProducts) error {
	// 使用 GORM 的 Create 方法将产品插入数据库
	return global.DBTour.Create(product).Error
}

// GetProductByID 根据 ID 查询产品
func GetProductByID(id int) (*model.TourProducts, error) {
	var product model.TourProducts
	if err := global.DBTour.First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// GetProductsByCategory 根据分类查询产品
func GetProductsByCategory(category string) ([]model.TourProducts, error) {
	var products []model.TourProducts
	if err := global.DBTour.Where("category = ?", category).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// CreateCategory 将分类保存到数据库
func CreateCategory(category *model.TourCategory) error {
	return global.DBTour.Create(category).Error
}
