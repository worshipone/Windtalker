// 自动生成模板TaskManage
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type TaskManage struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:参数;type:varchar(20);size:20;"`
      Value  string `json:"value" form:"value" gorm:"column:value;comment:值;type:varchar(20);size:20;"`
}


func (TaskManage) TableName() string {
  return "taskManage"
}

