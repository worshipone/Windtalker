// 自动生成模板Setting
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Setting struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:参数;type:varchar(20);size:20;"`
}


func (Setting) TableName() string {
  return "setting"
}

