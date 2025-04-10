package service

import (
	"backend-svc-go/dao"
	"backend-svc-go/global"
	"backend-svc-go/model"
	"time"
)

func RegisterCompany(name, contact_name, phone, address string) (*model.Company, error) {
	var company model.Company
	company.Name = name
	company.Contact_name = contact_name
	company.Phone = phone
	company.Address = address
	company_no := GenCompanyNo()
	company.Company_no = company_no
	company.Created_at = time.Now()
	company.Updated_at = time.Now()
	return dao.RegisterCompany(&company)
}

func RegisterStore(company_no, name, contact_name, phone, address string) (*model.Store, error) {
	var store model.Store
	store.Name = name
	store.Contact_name = contact_name
	store.Phone = phone
	store.Address = address
	store.Company_no = company_no
	store.Created_at = time.Now()
	store.Updated_at = time.Now()
	return dao.RegisterStore(&store)
}

func CreateNote(store_id int, name string) (*model.Notes, error) {
	var note model.Notes
	note.Name = name
	note.Store_id = store_id
	note.Status = 0
	note.Created_at = time.Now()
	note.Updated_at = time.Now()
	return dao.CreateNote(&note)
}

func GenCompanyNo() string {
	return global.RandomString(10)
}
