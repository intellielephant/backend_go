package model

import "time"

type Company struct {
	Id           int       `gorm:"primaryKey"`
	Name         string    `json:"name"`
	Contact_name string    `json:"contact_name"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	Company_no   string    `json:"company_no"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

type Store struct {
	Id           int       `gorm:"primaryKey"`
	Company_no   string    `json:"company_no"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	Contact_name string    `json:"contact_name"`
	Phone        string    `json:"phone"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

type Business struct {
	Id         int       `gorm:"primaryKey"`
	Business   string    `json:"business"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Notes struct {
	Id         int       `gorm:"primaryKey"`
	Store_id   int       `json:"store_id"`
	Name       string    `json:"name"`
	Status     int       `json:"status"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Goods struct {
	Id         int       `gorm:"primaryKey"`
	Store_id   int       `json:"store_id"`
	Note_id    int       `json:"note_id"`
	Title      string    `json:"title"`
	Status     int       `json:"status"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
