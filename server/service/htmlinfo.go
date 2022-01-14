package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateHtmlinfo
//@description: 创建Htmlinfo记录
//@param: htmlinfo model.Htmlinfo
//@return: err error

func CreateHtmlinfo(htmlinfo model.Htmlinfo) (err error) {
	err = global.GVA_DB.Create(&htmlinfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteHtmlinfo
//@description: 删除Htmlinfo记录
//@param: htmlinfo model.Htmlinfo
//@return: err error

func DeleteHtmlinfo(htmlinfo model.Htmlinfo) (err error) {
	err = global.GVA_DB.Delete(&htmlinfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteHtmlinfoByIds
//@description: 批量删除Htmlinfo记录
//@param: ids request.IdsReq
//@return: err error

func DeleteHtmlinfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Htmlinfo{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateHtmlinfo
//@description: 更新Htmlinfo记录
//@param: htmlinfo *model.Htmlinfo
//@return: err error

func UpdateHtmlinfo(htmlinfo model.Htmlinfo) (err error) {
	err = global.GVA_DB.Save(&htmlinfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetHtmlinfo
//@description: 根据id获取Htmlinfo记录
//@param: id uint
//@return: err error, htmlinfo model.Htmlinfo

func GetHtmlinfo(id uint) (err error, htmlinfo model.Htmlinfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&htmlinfo).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetHtmlinfoInfoList
//@description: 分页获取Htmlinfo记录
//@param: info request.HtmlinfoSearch
//@return: err error, list interface{}, total int64

func GetHtmlinfoInfoList(info request.HtmlinfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Htmlinfo{})
    var htmlinfos []model.Htmlinfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Url != "" {
        db = db.Where("`url` LIKE ?","%"+ info.Url+"%")
    }
    if info.Engine != "" {
        db = db.Where("`engine` = ?",info.Engine)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&htmlinfos).Error
	return err, htmlinfos, total
}