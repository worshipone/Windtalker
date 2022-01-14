// 自动生成模板Htmlinfo
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type Htmlinfo struct {
      global.GVA_MODEL
      A  string `json:"a" form:"a" gorm:"column:a;comment:;type:varchar(2048);size:2048;"`
      Url  string `json:"url" form:"url" gorm:"column:url;comment:;type:varchar(255);size:255;"`
      Engine  string `json:"engine" form:"engine" gorm:"column:engine;comment:;type:varchar(255);size:255;"`
      H1  string `json:"h1" form:"h1" gorm:"column:h1;comment:;type:varchar(255);size:255;"`
      H2  string `json:"h2" form:"h2" gorm:"column:h2;comment:;type:varchar(255);size:255;"`
      H3  string `json:"h3" form:"h3" gorm:"column:h3;comment:;type:varchar(255);size:255;"`
      Label  string `json:"label" form:"label" gorm:"column:label;comment:;type:varchar(255);size:255;"`
      Li  string `json:"li" form:"li" gorm:"column:li;comment:;type:varchar(255);size:255;"`
      MetaDescription  string `json:"metaDescription" form:"metaDescription" gorm:"column:meta_description;comment:;type:varchar(2048);size:2048;"`
      MetaKeywords  string `json:"metaKeywords" form:"metaKeywords" gorm:"column:meta_keywords;comment:;type:varchar(255);size:255;"`
      P  string `json:"p" form:"p" gorm:"column:p;comment:;type:varchar(255);size:255;"`
      Span  string `json:"span" form:"span" gorm:"column:span;comment:;type:varchar(255);size:255;"`
      Time  time.Time `json:"time" form:"time" gorm:"column:time;comment:;type:date;"`
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;type:varchar(255);size:255;"`
}


func (Htmlinfo) TableName() string {
  return "htmlinfo"
}

