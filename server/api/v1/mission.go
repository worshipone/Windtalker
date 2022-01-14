package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Mission
// @Summary 创建Mission
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Mission true "创建Mission"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mission/createMission [post]
func CreateMission(c *gin.Context) {
	var mission model.Mission
	_ = c.ShouldBindJSON(&mission)
	if err := service.CreateMission(mission); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Mission
// @Summary 删除Mission
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Mission true "删除Mission"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mission/deleteMission [delete]
func DeleteMission(c *gin.Context) {
	var mission model.Mission
	_ = c.ShouldBindJSON(&mission)
	if err := service.DeleteMission(mission); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Mission
// @Summary 批量删除Mission
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Mission"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /mission/deleteMissionByIds [delete]
func DeleteMissionByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteMissionByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Mission
// @Summary 更新Mission
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Mission true "更新Mission"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mission/updateMission [put]
func UpdateMission(c *gin.Context) {
	var mission model.Mission
	_ = c.ShouldBindJSON(&mission)
	if err := service.UpdateMission(mission); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Mission
// @Summary 用id查询Mission
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Mission true "用id查询Mission"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mission/findMission [get]
func FindMission(c *gin.Context) {
	var mission model.Mission
	_ = c.ShouldBindQuery(&mission)
	if err, remission := service.GetMission(mission.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remission": remission}, c)
	}
}

// @Tags Mission
// @Summary 分页获取Mission列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MissionSearch true "分页获取Mission列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mission/getMissionList [get]
func GetMissionList(c *gin.Context) {
	var pageInfo request.MissionSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetMissionInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
