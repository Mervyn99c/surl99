// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameURLMapping = "url_mapping"

// URLMapping 短链和长链的映射表
type URLMapping struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键" json:"id"`                          // 自增主键
	Surl          string    `gorm:"column:surl;not null;comment:短链" json:"surl"`                                             // 短链
	Lurl          string    `gorm:"column:lurl;not null;comment:长链" json:"lurl"`                                             // 长链
	EffectiveDays int32     `gorm:"column:effective_days;not null;comment:defualt:6天" json:"effective_days"`                 // defualt:6天
	CreateTime    time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"create_time"` // 记录创建时间
	UpdateTime    time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:记录更新时间" json:"update_time"` // 记录更新时间
	UpdateBy      time.Time `gorm:"column:update_by;not null;comment:记录更新人" json:"update_by"`                                // 记录更新人
	Deleted       string    `gorm:"column:deleted;not null;comment:删除标记" json:"deleted"`                                     // 删除标记
}

// TableName URLMapping's table name
func (*URLMapping) TableName() string {
	return TableNameURLMapping
}
