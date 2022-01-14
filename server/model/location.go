// 自动生成模板Location
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Location struct {
      global.GVA_MODEL
      Province  string `json:"province" form:"province" gorm:"column:province;comment:省份;type:varchar(191);size:191;"`
      Count  int `json:"count" form:"count" gorm:"column:count;comment:数目;type:int;size:20;"`
}


func (Location) TableName() string {
  return "location"
}

