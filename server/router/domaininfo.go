package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDomaininfoRouter(Router *gin.RouterGroup) {
	DomaininfoRouter := Router.Group("domaininfo").Use(middleware.OperationRecord())
	{
		DomaininfoRouter.POST("createDomaininfo", v1.CreateDomaininfo)   // 新建Domaininfo
		DomaininfoRouter.DELETE("deleteDomaininfo", v1.DeleteDomaininfo) // 删除Domaininfo
		DomaininfoRouter.DELETE("deleteDomaininfoByIds", v1.DeleteDomaininfoByIds) // 批量删除Domaininfo
		DomaininfoRouter.PUT("updateDomaininfo", v1.UpdateDomaininfo)    // 更新Domaininfo
		DomaininfoRouter.GET("findDomaininfo", v1.FindDomaininfo)        // 根据ID获取Domaininfo
		DomaininfoRouter.GET("getDomaininfoList", v1.GetDomaininfoList)  // 获取Domaininfo列表
	}
}
