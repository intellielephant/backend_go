package service

import (
	"backend-svc-go/dao"
	"backend-svc-go/model"
	"time"
)

func AddGoods(name string, category int, price float64, image_url string) (*model.Goods, error) {
	var goods model.Goods
	goods.Name = name
	goods.Category = category
	goods.Price = price
	goods.Image_url = image_url
	goods.Status = 0
	goods.Created_at = time.Now()
	goods.Updated_at = time.Now()
	return dao.AddGoods(&goods)
}

func GetGoods() ([]*model.GoodsWithCategory, error) {
	return dao.GetGoods()
}

func UpdateGoods(goods model.Goods) error {
	return dao.UpdateGoods(goods)
}

func AddOrder(table_name string, goods_list []map[string]interface{}) (*model.Order, []*model.OrderGoods, error) {
	var order model.Order
	order.TableName = table_name
	order.Status = 0
	order.Created_at = time.Now()
	order.Updated_at = time.Now()
	return dao.AddOrder(&order, goods_list)

}

func GetOrderByTableName(table_name string) (*model.Order, []*model.OrderGoods, error) {
	return dao.GetOrderByTableName(table_name)

}

func GetGoodsCategory() ([]*model.Category, error) {
	return dao.GetGoodsCategory()
}

func GetOrders(start_date time.Time, end_date time.Time) ([]*model.Order, error) {
	return dao.GetOrders(start_date, end_date)
}

func MapToGoods(data map[string]interface{}) (model.Goods, error) {
	var goods model.Goods

	// 手动映射字段
	if value, ok := data["Id"]; ok {
		goods.Id = int(value.(float64)) // 假设 id 是 float64 类型
	}
	if value, ok := data["name"]; ok {
		goods.Name = value.(string)
	}
	if value, ok := data["category"]; ok {
		goods.Category = int(value.(float64)) // 假设 category 是 float64 类型
	}
	if value, ok := data["price"]; ok {
		goods.Price = value.(float64)
	}
	if value, ok := data["image_url"]; ok {
		goods.Image_url = value.(string)
	}
	if value, ok := data["status"]; ok {
		goods.Status = int(value.(float64)) // 假设 status 是 float64 类型
	}

	return goods, nil
}
