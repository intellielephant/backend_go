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
