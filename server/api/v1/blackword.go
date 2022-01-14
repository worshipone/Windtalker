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

// @Tags Blackword
// @Summary 创建Blackword
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Blackword true "创建Blackword"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /blackword/createBlackword [post]
func CreateBlackword(c *gin.Context) {
	var blackword model.Blackword
	_ = c.ShouldBindJSON(&blackword)
	if err := service.CreateBlackword(blackword); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Blackword
// @Summary 删除Blackword
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Blackword true "删除Blackword"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /blackword/deleteBlackword [delete]
func DeleteBlackword(c *gin.Context) {
	var blackword model.Blackword
	_ = c.ShouldBindJSON(&blackword)
	if err := service.DeleteBlackword(blackword); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Blackword
// @Summary 批量删除Blackword
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Blackword"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /blackword/deleteBlackwordByIds [delete]
func DeleteBlackwordByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteBlackwordByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Blackword
// @Summary 更新Blackword
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Blackword true "更新Blackword"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /blackword/updateBlackword [put]
func UpdateBlackword(c *gin.Context) {
	var blackword model.Blackword
	_ = c.ShouldBindJSON(&blackword)
	if err := service.UpdateBlackword(blackword); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Blackword
// @Summary 用id查询Blackword
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Blackword true "用id查询Blackword"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /blackword/findBlackword [get]
func FindBlackword(c *gin.Context) {
	var blackword model.Blackword
	_ = c.ShouldBindQuery(&blackword)
	if err, reblackword := service.GetBlackword(blackword.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reblackword": reblackword}, c)
	}
}

// @Tags Blackword
// @Summary 分页获取Blackword列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BlackwordSearch true "分页获取Blackword列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /blackword/getBlackwordList [get]
func GetBlackwordList(c *gin.Context) {
	var pageInfo request.BlackwordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetBlackwordInfoList(pageInfo); err != nil {
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
