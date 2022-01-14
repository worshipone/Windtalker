package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDomaincheck
//@description: 创建Domaincheck记录
//@param: domaincheck model.Domaincheck
//@return: err error

func CreateDomaincheck(domaincheck model.Domaincheck) (err error) {
	err = global.GVA_DB.Create(&domaincheck).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDomaincheck
//@description: 删除Domaincheck记录
//@param: domaincheck model.Domaincheck
//@return: err error

func DeleteDomaincheck(domaincheck model.Domaincheck) (err error) {
	err = global.GVA_DB.Delete(&domaincheck).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDomaincheckByIds
//@description: 批量删除Domaincheck记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDomaincheckByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Domaincheck{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDomaincheck
//@description: 更新Domaincheck记录
//@param: domaincheck *model.Domaincheck
//@return: err error

func UpdateDomaincheck(domaincheck model.Domaincheck) (err error) {
	err = global.GVA_DB.Save(&domaincheck).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDomaincheck
//@description: 根据id获取Domaincheck记录
//@param: id uint
//@return: err error, domaincheck model.Domaincheck

func GetDomaincheck(id uint) (err error, domaincheck model.Domaincheck) {
	err = global.GVA_DB.Where("id = ?", id).First(&domaincheck).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDomaincheckInfoList
//@description: 分页获取Domaincheck记录
//@param: info request.DomaincheckSearch
//@return: err error, list interface{}, total int64

func GetDomaincheckInfoList(info request.DomaincheckSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Domaincheck{})
    var domainchecks []model.Domaincheck
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&domainchecks).Error
	return err, domainchecks, total
}