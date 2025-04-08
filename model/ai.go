package model

import "time"

type Function struct {
	Id          int       `gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Prompt      string    `json:"prompt"`
	Category    string    `json:"category"`
	Count       int       `json:"count"`
	Created_at  time.Time `json:"created_at"`
}

type UserLogin struct {
	Id         int       `gorm:"primaryKey"`
	Name       string    `json:"name"`
	LoginType  string    `json:"login_type"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type BaiduAccessToken struct {
	ID          int    `gorm:"primaryKey"`
	AccessToken string `json:"access_token"`
	Platform    string `json:"platform"`
	ExpiresIn   int    `json:"expires_in"`
}

type App struct {
	Id          int       `gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Pic         string    `json:"pic"`
	Route       string    `json:"route"`
	CategoryID  int       `json:"category_id"`
	Created_at  time.Time `json:"created_at"`
}

type UserPhone struct {
	Id         int       `gorm:"primaryKey"`
	Phone      string    `json:"phone"`
	Nickname   string    `json:"nickname"`
	Avatar     string    `json:"avatar"`
	Created_at time.Time `json:"created_at"`
}

type UserWeixin struct {
	Id         int       `gorm:"primaryKey"`
	Openid     string    `json:"openid"`
	Unionid    string    `json:"unionid"`
	Nickname   string    `json:"nickname"`
	Avatar     string    `json:"avatar"`
	Created_at time.Time `json:"created_at"`
}

type UserRelation struct {
	Id           int       `gorm:"primaryKey"`
	UserPhoneID  int       `json:"user_phone_id"`
	UserWeixinID int       `json:"user_weixin_id"`
	InvitorCode  string    `json:"invitor_code"`
	InviteCode   string    `json:"invite_code"`
	Created_at   time.Time `json:"created_at"`
}

type User struct {
	Id          int        `gorm:"primaryKey"`
	UserPhone   UserPhone  `json:"user_phone"`
	UserWeixin  UserWeixin `json:"user_weixin"`
	InvitorCode string     `json:"invitor_code"`
	InviteCode  string     `json:"invite_code"`
}

type Predict struct {
	Money string `json:"money"`

	Text string `json:"text"`
}
