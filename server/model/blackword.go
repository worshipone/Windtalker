// 自动生成模板Blackword
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type Blackword struct {
      global.GVA_MODEL
      Seedword  string `json:"seedword" form:"seedword" gorm:"column:seedword;comment:;type:varchar(255);size:255;"`
      Time  time.Time `json:"time" form:"time" gorm:"column:time;comment:;type:datetime;"`
}


func (Blackword) TableName() string {
  return "blackword"
}

