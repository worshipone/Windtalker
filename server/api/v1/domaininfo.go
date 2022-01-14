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

// @Tags Domaininfo
// @Summary 创建Domaininfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Domaininfo true "创建Domaininfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /domaininfo/createDomaininfo [post]
func CreateDomaininfo(c *gin.Context) {
	var domaininfo model.Domaininfo
	_ = c.ShouldBindJSON(&domaininfo)
	if err := service.CreateDomaininfo(domaininfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Domaininfo
// @Summary 删除Domaininfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Domaininfo true "删除Domaininfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /domaininfo/deleteDomaininfo [delete]
func DeleteDomaininfo(c *gin.Context) {
	var domaininfo model.Domaininfo
	_ = c.ShouldBindJSON(&domaininfo)
	if err := service.DeleteDomaininfo(domaininfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Domaininfo
// @Summary 批量删除Domaininfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Domaininfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /domaininfo/deleteDomaininfoByIds [delete]
func DeleteDomaininfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDomaininfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Domaininfo
// @Summary 更新Domaininfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Domaininfo true "更新Domaininfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /domaininfo/updateDomaininfo [put]
func UpdateDomaininfo(c *gin.Context) {
	var domaininfo model.Domaininfo
	_ = c.ShouldBindJSON(&domaininfo)
	if err := service.UpdateDomaininfo(domaininfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Domaininfo
// @Summary 用id查询Domaininfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Domaininfo true "用id查询Domaininfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /domaininfo/findDomaininfo [get]
func FindDomaininfo(c *gin.Context) {
	var domaininfo model.Domaininfo
	_ = c.ShouldBindQuery(&domaininfo)
	if err, redomaininfo := service.GetDomaininfo(domaininfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redomaininfo": redomaininfo}, c)
	}
}

// @Tags Domaininfo
// @Summary 分页获取Domaininfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DomaininfoSearch true "分页获取Domaininfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /domaininfo/getDomaininfoList [get]
func GetDomaininfoList(c *gin.Context) {
	var pageInfo request.DomaininfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDomaininfoInfoList(pageInfo); err != nil {
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
