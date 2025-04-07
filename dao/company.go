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
