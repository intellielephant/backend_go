package dao

import (
	"backend-svc-go/global"
	"backend-svc-go/model"
)

func AddGoods(goods *model.FishGoods) (*model.FishGoods, error) {
	tx := global.DBLittleFish.Create(&goods)
	return goods, tx.Error
}
