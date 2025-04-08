package dao

import (
	"backend-svc-go/global"
	"backend-svc-go/model"
	"time"
)

func AddGoods(goods *model.Goods) (*model.Goods, error) {
	tx := global.DBLittleFish.Create(&goods)
	return goods, tx.Error
}

func GetGoods() ([]*model.GoodsWithCategory, error) {
	var goodsWithCategory []*model.GoodsWithCategory
	tx := global.DBLittleFish.Table("goods").
		Select("goods.id, goods.name, goods.status, goods.category, goods.price, goods.image_url, category.name as category_name").
		Joins("left join category on category.id = goods.category").
		Where("goods.status = ?", 0).Find(&goodsWithCategory)

	return goodsWithCategory, tx.Error
}

func AddOrder(order *model.Order, goods_list []map[string]interface{}) (*model.Order, []*model.OrderGoods, error) {
	tx := global.DBLittleFish.Create(&order)

	if tx.Error != nil {
		return order, nil, tx.Error
	}

	var order_goods []*model.OrderGoods
	for _, goods := range goods_list {
		var order_good model.OrderGoods
		order_good.OrderId = order.Id
		order_good.GoodsId = int(goods["goods_id"].(float64))
		order_good.GoodsName = goods["goods_name"].(string)
		order_good.Number = int(goods["number"].(float64))
		order_good.Price = goods["price"].(float64)
		order_good.Created_at = time.Now()
		order_good.Updated_at = time.Now()
		order_goods = append(order_goods, &order_good)
		tx := global.DBLittleFish.Create(&order_good)
		if tx.Error != nil {
			return order, order_goods, tx.Error
		}
	}

	return order, order_goods, tx.Error
}
