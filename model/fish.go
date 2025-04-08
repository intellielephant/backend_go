/*
 * @Author: Will
 * @Date: 2022-11-14 19:14:48
 * @LastEditors: Will
 * @LastEditTime: 2022-11-14 19:14:48
 * @Description: 请填写简介
 */
package model

import "time"

type Goods struct {
	Id         int       `gorm:"primaryKey"`
	Name       string    `json:"name"`
	Status     int       `json:"status"`
	Category   int       `json:"category"`
	Price      float64   `json:"price"`
	Image_url  string    `json:"image_url"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type GoodsWithCategory struct {
	Goods
	CategoryName string `json:"category_name"`
}

type Order struct {
	Id         int       `gorm:"primaryKey"`
	TableName  string    `json:"table_name"`
	Status     int       `json:"status"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Category struct {
	Id    int    `gorm:"primaryKey"`
	Name  string `json:"name"`
	Index int    `json:"index"`
}

type OrderGoods struct {
	Id         int       `gorm:"primaryKey"`
	OrderId    int       `json:"order_id"`
	GoodsId    int       `json:"goods_id"`
	Number     int       `json:"number"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
