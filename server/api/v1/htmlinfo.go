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

// @Tags Htmlinfo
// @Summary 创建Htmlinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Htmlinfo true "创建Htmlinfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /htmlinfo/createHtmlinfo [post]
func CreateHtmlinfo(c *gin.Context) {
	var htmlinfo model.Htmlinfo
	_ = c.ShouldBindJSON(&htmlinfo)
	if err := service.CreateHtmlinfo(htmlinfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Htmlinfo
// @Summary 删除Htmlinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Htmlinfo true "删除Htmlinfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /htmlinfo/deleteHtmlinfo [delete]
func DeleteHtmlinfo(c *gin.Context) {
	var htmlinfo model.Htmlinfo
	_ = c.ShouldBindJSON(&htmlinfo)
	if err := service.DeleteHtmlinfo(htmlinfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Htmlinfo
// @Summary 批量删除Htmlinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Htmlinfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /htmlinfo/deleteHtmlinfoByIds [delete]
func DeleteHtmlinfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteHtmlinfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Htmlinfo
// @Summary 更新Htmlinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Htmlinfo true "更新Htmlinfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /htmlinfo/updateHtmlinfo [put]
func UpdateHtmlinfo(c *gin.Context) {
	var htmlinfo model.Htmlinfo
	_ = c.ShouldBindJSON(&htmlinfo)
	if err := service.UpdateHtmlinfo(htmlinfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Htmlinfo
// @Summary 用id查询Htmlinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Htmlinfo true "用id查询Htmlinfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /htmlinfo/findHtmlinfo [get]
func FindHtmlinfo(c *gin.Context) {
	var htmlinfo model.Htmlinfo
	_ = c.ShouldBindQuery(&htmlinfo)
	if err, rehtmlinfo := service.GetHtmlinfo(htmlinfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehtmlinfo": rehtmlinfo}, c)
	}
}

// @Tags Htmlinfo
// @Summary 分页获取Htmlinfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.HtmlinfoSearch true "分页获取Htmlinfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /htmlinfo/getHtmlinfoList [get]
func GetHtmlinfoList(c *gin.Context) {
	var pageInfo request.HtmlinfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetHtmlinfoInfoList(pageInfo); err != nil {
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
