package dao

import (
	"backend-svc-go/global"
	"backend-svc-go/model"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

func GetFunction(category string) ([]*model.Function, error) {
	var functions []*model.Function
	var tx *gorm.DB
	if category == "" {
		tx = global.DBLittleIn.Table("function").Scan(&functions)
	} else {
		tx = global.DBAI.Table("function").Where("category = ?", category).Scan(&functions)
	}

	return functions, tx.Error
}

func QuickLogin(name, login_type string) (*model.UserLogin, error) {
	var user_login model.UserLogin
	tx := global.DBLittleIn.Where(&model.UserLogin{Name: name, LoginType: login_type}).FirstOrCreate(&user_login)
	return &user_login, tx.Error
}

func GetBaiduAccessToken() (*model.BaiduAccessToken, error) {
	var access_token model.BaiduAccessToken
	tx := global.DBAI.Table("access_token").Where("platform = ?", "baidu").First(&access_token)
	return &access_token, tx.Error
}

func GetCategory() ([]*model.Category, error) {
	var category []*model.Category
	tx := global.DBAI.Table("category").Find(&category)
	return category, tx.Error
}

func GetAppByCategoryID(category_id int) ([]*model.App, error) {
	var app []*model.App
	tx := global.DBAI.Table("app").Joins("JOIN category ON app.category_id = category.id").Where("category.id = ?", category_id).Find(&app)
	return app, tx.Error
}

func GetHotApp() ([]*model.App, error) {
	var app []*model.App
	tx := global.DBAI.Table("app").Joins("JOIN category ON app.category_id = category.id").Where("category.name = ?", "热门").Find(&app)
	return app, tx.Error
}

func UserPhoneLogin(phone, invitor_code string) (*model.User, error) {
	user := model.User{}
	userPhone := model.UserPhone{}
	tx := global.DBAI.Table("user_phone").Where("phone=?", phone).First(&userPhone)
	userRelation := model.UserRelation{}
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		userPhone.Phone = phone
		userPhone.Created_at = time.Now()
		tx = global.DBAI.Table("user_phone").Where("phone=?", phone).Create(&userPhone)
		if tx.Error == nil {
			user.UserPhone = userPhone
			userRelation.UserPhoneID = userPhone.Id
			userRelation.InvitorCode = invitor_code
			user.InvitorCode = invitor_code
			invite_code := global.GenInviteCode()
			user.InviteCode = invite_code
			userRelation.InviteCode = invite_code
			userRelation.Created_at = time.Now()
			tx = global.DBAI.Table("user_relation").Create(&userRelation)
			if tx.Error == nil {
				user.Id = userRelation.Id
			}
			return &user, tx.Error
		}
	}

	tx = global.DBAI.Table("user_relation").Where("user_phone_id = ?", userPhone.Id).First(&userRelation)
	if tx.Error == nil {
		user.Id = userRelation.Id
		user.UserPhone = userPhone
		user.InvitorCode = userRelation.InvitorCode
		user.InviteCode = userRelation.InviteCode
	}
	return &user, tx.Error
}

func CreateUserRelation(user_phone_id, user_weixin_id int, invite_code, invitor_code string) (*model.UserRelation, error) {
	var userRelation = model.UserRelation{}
	userRelation.UserPhoneID = user_phone_id
	userRelation.UserWeixinID = user_weixin_id
	userRelation.InviteCode = invite_code
	userRelation.InvitorCode = invitor_code
	userRelation.Created_at = time.Now()
	tx := global.DBAI.Table("user_relation").Create(&userRelation)
	return &userRelation, tx.Error
}

func GetUserByOpenid(openid, accessToken, invitorCode string) (*model.User, error) {
	var user model.User
	var userWeixin model.UserWeixin
	var userPhone model.UserPhone
	var userRelation model.UserRelation
	tx := global.DBAI.Table("user_weixin").Where("openid = ?", openid).First(&userWeixin)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		userinfoRes, err := global.GetUserinfoByAccessToken(accessToken, openid)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		userWeixin, err := CreateWeixinUser(openid, userinfoRes.Get("unionid").MustString(), userinfoRes.Get("nickname").MustString(), userinfoRes.Get("headimgurl").MustString())
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		invite_code := global.GenInviteCode()
		userRelation, err := CreateUserRelation(0, userWeixin.Id, invite_code, invitorCode)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		user.Id = userRelation.Id
		user.UserWeixin = *userWeixin
		user.UserPhone = model.UserPhone{}
		user.InvitorCode = invitorCode
		user.InviteCode = invite_code

	} else {
		user.UserWeixin = userWeixin
		tx = global.DBAI.Table("user_relation").Where("user_weixin_id = ?", userWeixin.Id).First(&userRelation)

		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			log.Fatal(tx.Error)
			return nil, tx.Error
		}
		if userRelation.UserPhoneID != 0 {
			tx = global.DBAI.Table("user_phone").Where("id = ?", userRelation.UserPhoneID).First(&userPhone)
			user.UserPhone = userPhone
		} else {
			user.UserPhone = model.UserPhone{}
		}
		user.Id = userRelation.Id
		user.InviteCode = userRelation.InviteCode
		user.InvitorCode = userRelation.InvitorCode

	}

	return &user, tx.Error
}

func GetUserByInvitor(invitor_code string) ([]model.UserRelation, error) {
	var userRelation = []model.UserRelation{}
	tx := global.DBAI.Table("user_relation").Where("invitor_code=?", invitor_code).Find(&userRelation)
	return userRelation, tx.Error
}

func GetUserWeixinInvitors(invitor, invitorType int) ([]model.UserWeixin, error) {
	var userWeixin = []model.UserWeixin{}
	tx := global.DBAI.Table("user_weixin").Where("invitor=? AND invitor_type=?", invitor, invitorType).Find(&userWeixin)
	return userWeixin, tx.Error
}

func GetUserWeixinByOpenid(openid string) (*model.UserWeixin, error) {
	userWeixin := model.UserWeixin{}
	tx := global.DBAI.Table("user_weixin").Where("openid=?", openid).First(&userWeixin)

	return &userWeixin, tx.Error
}

func CreateWeixinUser(openid, unionid, nickname, avatar string) (*model.UserWeixin, error) {
	userWeixin := model.UserWeixin{}
	userWeixin.Openid = openid
	userWeixin.Unionid = unionid
	userWeixin.Avatar = avatar
	userWeixin.Nickname = nickname
	userWeixin.Created_at = time.Now()
	tx := global.DBAI.Table("user_weixin").Create(&userWeixin)
	return &userWeixin, tx.Error
}

func UpdateAvatar(user_id int, avatar string) error {
	user := model.UserRelation{}
	tx := global.DBAI.Table("user_relation").Where("id = ?", user_id).First(&user)
	if tx.Error != nil {
		return tx.Error
	} else {
		tx = global.DBAI.Table("user_phone").Where("id = ?", user.UserPhoneID).Update("avatar", avatar)
		return tx.Error
	}
}

func IsOpenidExist(openid string) bool {
	var userWeixin model.UserWeixin
	tx := global.DBAI.Table("user_weixin").Where("openid=?", openid).First(&userWeixin)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return false
	} else {
		return true
	}
}
