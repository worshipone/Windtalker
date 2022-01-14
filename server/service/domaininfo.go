package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDomaininfo
//@description: 创建Domaininfo记录
//@param: domaininfo model.Domaininfo
//@return: err error

func CreateDomaininfo(domaininfo model.Domaininfo) (err error) {
	err = global.GVA_DB.Create(&domaininfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDomaininfo
//@description: 删除Domaininfo记录
//@param: domaininfo model.Domaininfo
//@return: err error

func DeleteDomaininfo(domaininfo model.Domaininfo) (err error) {
	err = global.GVA_DB.Delete(&domaininfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDomaininfoByIds
//@description: 批量删除Domaininfo记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDomaininfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Domaininfo{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDomaininfo
//@description: 更新Domaininfo记录
//@param: domaininfo *model.Domaininfo
//@return: err error

func UpdateDomaininfo(domaininfo model.Domaininfo) (err error) {
	err = global.GVA_DB.Save(&domaininfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDomaininfo
//@description: 根据id获取Domaininfo记录
//@param: id uint
//@return: err error, domaininfo model.Domaininfo

func GetDomaininfo(id uint) (err error, domaininfo model.Domaininfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&domaininfo).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDomaininfoInfoList
//@description: 分页获取Domaininfo记录
//@param: info request.DomaininfoSearch
//@return: err error, list interface{}, total int64

func GetDomaininfoInfoList(info request.DomaininfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Domaininfo{})
    var domaininfos []model.Domaininfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.DomainName != "" {
        db = db.Where("`domain_name` LIKE ?","%"+ info.DomainName+"%")
    }
    if info.Ip != "" {
        db = db.Where("`ip` = ?",info.Ip)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&domaininfos).Error
	return err, domaininfos, total
}