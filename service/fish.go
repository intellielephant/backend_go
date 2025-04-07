package service

import (
	"backend-svc-go/dao"
	"backend-svc-go/model"
	"time"
)

func AddGoods(name string, category int, price float64, image_url string) (*model.FishGoods, error) {
	var goods model.FishGoods
	goods.Name = name
	goods.Category = category
	goods.Price = price
	goods.Image_url = image_url
	goods.Status = 0
	goods.Created_at = time.Now()
	goods.Updated_at = time.Now()
	return dao.AddGoods(&goods)
}
