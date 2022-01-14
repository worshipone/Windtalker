package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDomaincheckRouter(Router *gin.RouterGroup) {
	DomaincheckRouter := Router.Group("domaincheck").Use(middleware.OperationRecord())
	{
		DomaincheckRouter.POST("createDomaincheck", v1.CreateDomaincheck)   // 新建Domaincheck
		DomaincheckRouter.DELETE("deleteDomaincheck", v1.DeleteDomaincheck) // 删除Domaincheck
		DomaincheckRouter.DELETE("deleteDomaincheckByIds", v1.DeleteDomaincheckByIds) // 批量删除Domaincheck
		DomaincheckRouter.PUT("updateDomaincheck", v1.UpdateDomaincheck)    // 更新Domaincheck
		DomaincheckRouter.GET("findDomaincheck", v1.FindDomaincheck)        // 根据ID获取Domaincheck
		DomaincheckRouter.GET("getDomaincheckList", v1.GetDomaincheckList)  // 获取Domaincheck列表
	}
}
