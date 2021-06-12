package model

import "time"

type Model struct {
	ID        uint      `json:"id" gorm:"primary_key;comment:'primary key'"`
	CreatedBy string    `json:"created_by" gorm:"comment:'creater'"`
	UpdatedBy string    `json:"modified_by" gorm:"comment:'updater'"`
	CreatedAt time.Time `json:"created_on" gorm:"comment:'create time'"`
	UpdatedAt time.Time `json:"modified_on" gorm:"comment:'update time'"`
	DeletedAt time.Time `json:"deleted_on" gorm:"comment:'delete time'"`
	IsDel     uint8     `json:"is_del" gorm:"comment:'soft delete'"`
}
