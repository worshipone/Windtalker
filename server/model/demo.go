// 自动生成模板Demo
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Demo struct {
      global.GVA_MODEL
      Province  string `json:"province" form:"province" gorm:"column:province;comment:演示;type:varchar(191);size:191;"`
      Count  int `json:"count" form:"count" gorm:"column:count;comment:数目;type:int;size:20;"`
}


func (Demo) TableName() string {
  return "demo"
}

