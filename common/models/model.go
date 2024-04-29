package models

import "time"

type Model struct {
	ID        int64     `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt;not null;autoCreateTime;comment:'创建时间'" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;not null;autoUpdateTime;comment:'更新时间'" json:"updatedAt"`
	DeletedAt time.Time `gorm:"column:deletedAt;index;comment:'删除时间'" json:"deletedAt"`
}
