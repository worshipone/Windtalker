// 自动生成模板Domaincheck
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type Domaincheck struct {
      global.GVA_MODEL
      Number  int `json:"number" form:"number" gorm:"column:number;comment:;type:int;size:10;"`
      DomainName  string `json:"domainName" form:"domainName" gorm:"column:domain_name;comment:;type:varchar(50);size:50;"`
      Ip  string `json:"ip" form:"ip" gorm:"column:ip;comment:;type:varchar(20);size:20;"`
      Dnssec  string `json:"dnssec" form:"dnssec" gorm:"column:dnssec;comment:;type:varchar(255);size:255;"`
      Created  time.Time `json:"created" form:"created" gorm:"column:created;comment:;type:datetime;"`
      Expires  time.Time `json:"expires" form:"expires" gorm:"column:expires;comment:;type:datetime;"`
      NameServers  string `json:"nameServers" form:"nameServers" gorm:"column:name_servers;comment:;type:varchar(255);size:255;"`
      RegistrantAddress  string `json:"registrantAddress" form:"registrantAddress" gorm:"column:registrant_address;comment:;type:varchar(255);size:255;"`
      RegistrantCity  string `json:"registrantCity" form:"registrantCity" gorm:"column:registrant_city;comment:;type:varchar(255);size:255;"`
      RegistrantCountry  string `json:"registrantCountry" form:"registrantCountry" gorm:"column:registrant_country;comment:;type:varchar(255);size:255;"`
      RegistrantName  string `json:"registrantName" form:"registrantName" gorm:"column:registrant_name;comment:;type:varchar(255);size:255;"`
      RegistrantOrganization  string `json:"registrantOrganization" form:"registrantOrganization" gorm:"column:registrant_organization;comment:;type:varchar(255);size:255;"`
      RegistrantState  string `json:"registrantState" form:"registrantState" gorm:"column:registrant_state;comment:;type:varchar(255);size:255;"`
      RegistrantZipcode  string `json:"registrantZipcode" form:"registrantZipcode" gorm:"column:registrant_zipcode;comment:;type:varchar(255);size:255;"`
      Registrar  string `json:"registrar" form:"registrar" gorm:"column:registrar;comment:;type:varchar(255);size:255;"`
      Updated  time.Time `json:"updated" form:"updated" gorm:"column:updated;comment:;type:datetime;"`
}


func (Domaincheck) TableName() string {
  return "domaincheck"
}

