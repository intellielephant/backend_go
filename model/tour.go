package model

import "time"

type TourProducts struct {
	Id                 int       `gorm:"primaryKey" json:"id"`             // 主键
	Name               string    `json:"name"`                             // 商品名称
	Description        string    `json:"description"`                      // 商品描述
	Status             int       `json:"status"`                           // 商品状态 (例如: 0-下架, 1-上架)
	Category           string    `json:"category"`                         // 分类
	Price              float64   `json:"price"`                            // 商品价格
	Duration           int       `json:"duration"`                         // 行程天数
	ImageUrl           string    `json:"image_url"`                        // 商品图片 URL
	Departure_date     time.Time `json:"departure_date"`                   // 出发日期
	Return_date        time.Time `json:"return_date"`                      // 返回日期
	Departure_location string    `json:"departure_location"`               // 出发地点
	Destination        string    `json:"destination"`                      // 目的地
	Availability       int       `json:"availability"`                     // 库存(可用人数)
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"` // 创建时间
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"` // 更新时间
}

type TourCategory struct {
	Id          int       `gorm:"primaryKey" json:"id"`             // 主键
	Name        string    `json:"name"`                             // 分类名称
	Description string    `json:"description"`                      // 分类描述
	ParentId    int       `json:"parent_id"`                        // 父分类 ID (如果是顶级分类则为 0)
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"` // 创建时间
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"` // 更新时间
}

type Users struct {
	Id         int       `gorm:"primaryKey" json:"id"`                              // 主键
	Username   string    `json:"username"`                                          // 用户名
	Email      string    `json:"email"`                                             // 邮箱
	Password   string    `json:"password"`                                          // 密码 (加密存储)
	Phone      string    `json:"phone"`                                             // 手机号码
	Wechat_id  string    `json:"wechat_id"`                                         // 微信 ID
	Avatar_url string    `json:"avatar_url"`                                        // 头像 URL
	Gender     string    `gorm:"type:enum('male','female', 'other')" json:"gender"` // 性别
	Birth_date time.Time `json:"birth_date"`                                        // 出生日期
	Address    string    `json:"address"`                                           // 地址
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`                  // 创建时间
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`                  // 更新时间
}
