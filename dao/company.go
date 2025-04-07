package dao

import (
	"backend-svc-go/global"
	"backend-svc-go/model"
)

func RegisterCompany(company *model.Company) (*model.Company, error) {
	tx := global.DBLittleIn.Create(&company)

	return company, tx.Error
}

func RegisterStore(store *model.Store) (*model.Store, error) {
	tx := global.DBLittleIn.Create(&store)

	return store, tx.Error
}

func CreateNote(note *model.Notes) (*model.Notes, error) {
	tx := global.DBLittleIn.Create(&note)

	return note, tx.Error
}

func CreateGoods(goods *model.Goods) (*model.Goods, error) {
	tx := global.DBLittleIn.Create(&goods)

	return goods, tx.Error
}

func SelectGoods(page_size, page_index int) (*[]model.Goods, error) {
	var goods []model.Goods
	tx := global.DBLittleIn.Table("goods").Limit(page_size).Offset((page_index-1)*page_size).Where("status = ?", 0).Scan(&goods)

	return &goods, tx.Error
}
