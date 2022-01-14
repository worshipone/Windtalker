// 自动生成模板Search
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Search struct {
      global.GVA_MODEL
      Url  string `json:"url" form:"url" gorm:"column:url;comment:网址;type:varchar(20);size:20;"`
}


func (Search) TableName() string {
  return "search"
}

