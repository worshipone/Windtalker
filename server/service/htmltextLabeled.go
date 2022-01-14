package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateHtmltextLabeled
//@description: 创建HtmltextLabeled记录
//@param: htmltextLabeled model.HtmltextLabeled
//@return: err error

func CreateHtmltextLabeled(htmltextLabeled model.HtmltextLabeled) (err error) {
	err = global.GVA_DB.Create(&htmltextLabeled).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteHtmltextLabeled
//@description: 删除HtmltextLabeled记录
//@param: htmltextLabeled model.HtmltextLabeled
//@return: err error

func DeleteHtmltextLabeled(htmltextLabeled model.HtmltextLabeled) (err error) {
	err = global.GVA_DB.Delete(&htmltextLabeled).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteHtmltextLabeledByIds
//@description: 批量删除HtmltextLabeled记录
//@param: ids request.IdsReq
//@return: err error

func DeleteHtmltextLabeledByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.HtmltextLabeled{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateHtmltextLabeled
//@description: 更新HtmltextLabeled记录
//@param: htmltextLabeled *model.HtmltextLabeled
//@return: err error

func UpdateHtmltextLabeled(htmltextLabeled model.HtmltextLabeled) (err error) {
	err = global.GVA_DB.Save(&htmltextLabeled).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetHtmltextLabeled
//@description: 根据id获取HtmltextLabeled记录
//@param: id uint
//@return: err error, htmltextLabeled model.HtmltextLabeled

func GetHtmltextLabeled(id uint) (err error, htmltextLabeled model.HtmltextLabeled) {
	err = global.GVA_DB.Where("id = ?", id).First(&htmltextLabeled).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetHtmltextLabeledInfoList
//@description: 分页获取HtmltextLabeled记录
//@param: info request.HtmltextLabeledSearch
//@return: err error, list interface{}, total int64

func GetHtmltextLabeledInfoList(info request.HtmltextLabeledSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.HtmltextLabeled{})
    var htmltextLabeleds []model.HtmltextLabeled
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&htmltextLabeleds).Error
	return err, htmltextLabeleds, total
}