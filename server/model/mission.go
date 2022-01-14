// 自动生成模板Mission
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Mission struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;type:varchar;"`
      Desc  string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;type:varchar;"`
      Seedword  string `json:"seedword" form:"seedword" gorm:"column:seedword;comment:种子词;type:varchar;"`
      Method  string `json:"method" form:"method" gorm:"column:method;comment:方式;type:varchar;"`
      Engine  string `json:"engine" form:"engine" gorm:"column:engine;comment:搜素引擎;type:char;"`
      Process  int `json:"process" form:"process" gorm:"column:process;comment:进度;type:varchar;"`
      White  string `json:"white" form:"white" gorm:"column:white;comment:白名单;type:varchar;"`
      Black  string `json:"black" form:"black" gorm:"column:black;comment:黑名单;type:varchar;"`
      DialogVisible  bool `json:"dialogVisible" form:"dialogVisible" gorm:"column:dialogVisible;comment:是否可见;type:varchar;"`
}


func (Mission) TableName() string {
  return "mission"
}

