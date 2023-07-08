package address

import "simple-mall/models"

// Address 地址管理表
type Address struct {
	models.BaseModel
	UserId          int32  `gorm:"type:int; not null; comment '用户id'" json:"userId"`
	Name            string `gorm:"type:varbinary(30); not null; comment '姓名'" json:"name"`
	Phone           string `gorm:"type:varbinary(20); not null; comment '手机号'" json:"phone"`
	Province        int    `gorm:"type:int; not null; comment '省code'" json:"province"`
	ProvinceName    string `gorm:"type:varbinary(50); not null; comment '省名称'" json:"provinceName"`
	City            int    `gorm:"type:int; not null; comment '市code'" json:"city"`
	CityName        string `gorm:"type:varbinary(50); not null; comment '市名称'" json:"cityName"`
	District        int    `gorm:"type:int; comment '区code'" json:"district"`
	DistrictName    string `gorm:"type:varbinary(50); comment '区名称'" json:"districtName"`
	Street          int    `gorm:"type:int; comment '街道code'" json:"street"`
	StreetName      string `gorm:"type:varbinary(50); comment '街道名称'" json:"streetName"`
	DetailedAddress string `gorm:"type:varbinary(500); not null; comment '详细地址'" json:"detailedAddress"`
	ZipCode         string `gorm:"type:varbinary(20); comment '邮政编码'" json:"zipCode"`
	DefaultAddress  string `gorm:"type:varbinary(10); default: 01; comment '默认地址 01 是 00 否'" json:"defaultAddress"`
}

// AreaInfo 区域信息
type AreaInfo struct {
	ID           int     `json:"id"`
	PID          int     `json:"pid"`
	Deep         int     `json:"deep"`
	Name         string  `json:"name"`
	PinyinPrefix string  `json:"pinyinPrefix"`
	Pinyin       string  `json:"pinyin"`
	ExtID        float64 `json:"extID"`
	ExtName      string  `json:"extName"`
}

// ListForm 分页地址信息
type ListForm struct {
	models.PaginationParameters
	ProvinceName string `json:"provinceName" form:"provinceName"`
	CityName     string `json:"cityName" form:"cityName"`
	DistrictName string `json:"districtName" form:"districtName"`
	StreetName   string `json:"streetName" form:"streetName"`
}

// SaveForm 地址保存表单
type SaveForm struct {
	ID              int32  `form:"id" json:"id"`
	UserId          int32  `json:"userId" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Phone           string `json:"phone" binding:"required"`
	Province        int    `json:"province" binding:"required"`
	ProvinceName    string `json:"provinceName" binding:"required"`
	City            int    `json:"city" binding:"required"`
	CityName        string `json:"cityName" binding:"required"`
	District        int    `json:"district"`
	DistrictName    string `json:"districtName"`
	Street          int    `json:"street"`
	StreetName      string `json:"streetName"`
	DetailedAddress string `json:"detailedAddress" binding:"required"`
	ZipCode         string `json:"zipCode"`
	DefaultAddress  string `json:"defaultAddress" binding:"required"`
}
