package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateBlackword
//@description: 创建Blackword记录
//@param: blackword model.Blackword
//@return: err error

func CreateBlackword(blackword model.Blackword) (err error) {
	err = global.GVA_DB.Create(&blackword).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBlackword
//@description: 删除Blackword记录
//@param: blackword model.Blackword
//@return: err error

func DeleteBlackword(blackword model.Blackword) (err error) {
	err = global.GVA_DB.Delete(&blackword).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBlackwordByIds
//@description: 批量删除Blackword记录
//@param: ids request.IdsReq
//@return: err error

func DeleteBlackwordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Blackword{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBlackword
//@description: 更新Blackword记录
//@param: blackword *model.Blackword
//@return: err error

func UpdateBlackword(blackword model.Blackword) (err error) {
	err = global.GVA_DB.Save(&blackword).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBlackword
//@description: 根据id获取Blackword记录
//@param: id uint
//@return: err error, blackword model.Blackword

func GetBlackword(id uint) (err error, blackword model.Blackword) {
	err = global.GVA_DB.Where("id = ?", id).First(&blackword).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBlackwordInfoList
//@description: 分页获取Blackword记录
//@param: info request.BlackwordSearch
//@return: err error, list interface{}, total int64

func GetBlackwordInfoList(info request.BlackwordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Blackword{})
    var blackwords []model.Blackword
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Seedword != "" {
        db = db.Where("`seedword` = ?",info.Seedword)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&blackwords).Error
	return err, blackwords, total
}