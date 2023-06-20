package models

import (
	"gorm.io/plugin/soft_delete"
	"simple-mall/utils/localTime"
)

// BaseModel 基础数据
type BaseModel struct {
	ID        int32                 `gorm:"primaryKey; auto_increment; comment '主键'" json:"id"`
	CreatedAt *localTime.LocalTime  `gorm:"column=created_at; comment '创建时间'" json:"createdAt"`
	UpdatedAt *localTime.LocalTime  `gorm:"column=updated_at; comment '更新时间'" json:"updatedAt"`
	DeletedAt *localTime.LocalTime  `gorm:"column=deleted_at; default:null; comment '删除时间'" json:"-"`
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt; comment '删除标志 0 1'" json:"-"`
}

// PaginationParameters 分页参数
type PaginationParameters struct {
	PageSize int `json:"pageSize" form:"pageSize" example:"10" binding:"required,min=0"`
	PageNo   int `json:"pageNo" form:"pageNo" example:"1" binding:"required,min=0"`
}

// PagingData 分页数据对象
type PagingData struct {
	List     any   `json:"list"`
	PageSize int   `json:"pageSize"`
	PageNo   int   `json:"pageNo"`
	Total    int32 `json:"total"`
}
