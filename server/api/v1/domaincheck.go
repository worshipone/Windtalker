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

// @Tags Domaincheck
// @Summary 创建Domaincheck
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Domaincheck true "创建Domaincheck"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /domaincheck/createDomaincheck [post]
func CreateDomaincheck(c *gin.Context) {
	var domaincheck model.Domaincheck
	_ = c.ShouldBindJSON(&domaincheck)
	if err := service.CreateDomaincheck(domaincheck); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Domaincheck
// @Summary 删除Domaincheck
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Domaincheck true "删除Domaincheck"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /domaincheck/deleteDomaincheck [delete]
func DeleteDomaincheck(c *gin.Context) {
	var domaincheck model.Domaincheck
	_ = c.ShouldBindJSON(&domaincheck)
	if err := service.DeleteDomaincheck(domaincheck); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Domaincheck
// @Summary 批量删除Domaincheck
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Domaincheck"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /domaincheck/deleteDomaincheckByIds [delete]
func DeleteDomaincheckByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDomaincheckByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Domaincheck
// @Summary 更新Domaincheck
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Domaincheck true "更新Domaincheck"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /domaincheck/updateDomaincheck [put]
func UpdateDomaincheck(c *gin.Context) {
	var domaincheck model.Domaincheck
	_ = c.ShouldBindJSON(&domaincheck)
	if err := service.UpdateDomaincheck(domaincheck); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Domaincheck
// @Summary 用id查询Domaincheck
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Domaincheck true "用id查询Domaincheck"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /domaincheck/findDomaincheck [get]
func FindDomaincheck(c *gin.Context) {
	var domaincheck model.Domaincheck
	_ = c.ShouldBindQuery(&domaincheck)
	if err, redomaincheck := service.GetDomaincheck(domaincheck.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redomaincheck": redomaincheck}, c)
	}
}

// @Tags Domaincheck
// @Summary 分页获取Domaincheck列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DomaincheckSearch true "分页获取Domaincheck列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /domaincheck/getDomaincheckList [get]
func GetDomaincheckList(c *gin.Context) {
	var pageInfo request.DomaincheckSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDomaincheckInfoList(pageInfo); err != nil {
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
