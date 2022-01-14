package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateMission
//@description: 创建Mission记录
//@param: mission model.Mission
//@return: err error

func CreateMission(mission model.Mission) (err error) {
	err = global.GVA_DB.Create(&mission).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMission
//@description: 删除Mission记录
//@param: mission model.Mission
//@return: err error

func DeleteMission(mission model.Mission) (err error) {
	err = global.GVA_DB.Delete(&mission).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMissionByIds
//@description: 批量删除Mission记录
//@param: ids request.IdsReq
//@return: err error

func DeleteMissionByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Mission{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateMission
//@description: 更新Mission记录
//@param: mission *model.Mission
//@return: err error

func UpdateMission(mission model.Mission) (err error) {
	err = global.GVA_DB.Save(&mission).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMission
//@description: 根据id获取Mission记录
//@param: id uint
//@return: err error, mission model.Mission

func GetMission(id uint) (err error, mission model.Mission) {
	err = global.GVA_DB.Where("id = ?", id).First(&mission).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMissionInfoList
//@description: 分页获取Mission记录
//@param: info request.MissionSearch
//@return: err error, list interface{}, total int64

func GetMissionInfoList(info request.MissionSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Mission{})
    var missions []model.Mission
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` = ?",info.Name)
    }
    if info.Desc != "" {
        db = db.Where("`desc` = ?",info.Desc)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&missions).Error
	return err, missions, total
}