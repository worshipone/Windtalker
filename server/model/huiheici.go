// 自动生成模板Huiheici
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Huiheici struct {
      global.GVA_MODEL
      Index  int `json:"index" form:"index" gorm:"column:index;comment:;type:int;size:10;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(45);size:45;"`
      Value  string `json:"value" form:"value" gorm:"column:value;comment:;type:varchar(45);size:45;"`
}


func (Huiheici) TableName() string {
  return "huiheici"
}

