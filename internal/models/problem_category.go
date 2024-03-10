package models

import (
	"gorm.io/gorm"
)

// problem_basic -> problem_category -> category_basic
// 其实我感觉不用problem_category也可以的，直接 problem_basic -> category_basic，感觉这样更加直接一点
type ProblemCategory struct {
	ID        uint           `gorm:"primarykey;" json:"id"`
	CreatedAt MyTime         `json:"created_at"`
	UpdatedAt MyTime         `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;" json:"deleted_at"`

	ProblemId     uint           `gorm:"column:problem_id;type:int(11);" json:"problem_id"`           // 问题的ID
	CategoryId    uint           `gorm:"column:category_id;type:int(11);" json:"category_id"`         // 分类的ID
	CategoryBasic *CategoryBasic `gorm:"foreignKey:id;references:category_id;" json:"category_basic"` // 关联分类的基础信息表
}

func (table *ProblemCategory) TableName() string {
	return "problem_category"
}
