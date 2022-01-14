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

// @Tags HtmltextLabeled
// @Summary 创建HtmltextLabeled
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HtmltextLabeled true "创建HtmltextLabeled"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /htmltextLabeled/createHtmltextLabeled [post]
func CreateHtmltextLabeled(c *gin.Context) {
	var htmltextLabeled model.HtmltextLabeled
	_ = c.ShouldBindJSON(&htmltextLabeled)
	if err := service.CreateHtmltextLabeled(htmltextLabeled); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags HtmltextLabeled
// @Summary 删除HtmltextLabeled
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HtmltextLabeled true "删除HtmltextLabeled"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /htmltextLabeled/deleteHtmltextLabeled [delete]
func DeleteHtmltextLabeled(c *gin.Context) {
	var htmltextLabeled model.HtmltextLabeled
	_ = c.ShouldBindJSON(&htmltextLabeled)
	if err := service.DeleteHtmltextLabeled(htmltextLabeled); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags HtmltextLabeled
// @Summary 批量删除HtmltextLabeled
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HtmltextLabeled"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /htmltextLabeled/deleteHtmltextLabeledByIds [delete]
func DeleteHtmltextLabeledByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteHtmltextLabeledByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags HtmltextLabeled
// @Summary 更新HtmltextLabeled
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HtmltextLabeled true "更新HtmltextLabeled"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /htmltextLabeled/updateHtmltextLabeled [put]
func UpdateHtmltextLabeled(c *gin.Context) {
	var htmltextLabeled model.HtmltextLabeled
	_ = c.ShouldBindJSON(&htmltextLabeled)
	if err := service.UpdateHtmltextLabeled(htmltextLabeled); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags HtmltextLabeled
// @Summary 用id查询HtmltextLabeled
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HtmltextLabeled true "用id查询HtmltextLabeled"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /htmltextLabeled/findHtmltextLabeled [get]
func FindHtmltextLabeled(c *gin.Context) {
	var htmltextLabeled model.HtmltextLabeled
	_ = c.ShouldBindQuery(&htmltextLabeled)
	if err, rehtmltextLabeled := service.GetHtmltextLabeled(htmltextLabeled.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehtmltextLabeled": rehtmltextLabeled}, c)
	}
}

// @Tags HtmltextLabeled
// @Summary 分页获取HtmltextLabeled列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.HtmltextLabeledSearch true "分页获取HtmltextLabeled列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /htmltextLabeled/getHtmltextLabeledList [get]
func GetHtmltextLabeledList(c *gin.Context) {
	var pageInfo request.HtmltextLabeledSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetHtmltextLabeledInfoList(pageInfo); err != nil {
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
