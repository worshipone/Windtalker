package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitHtmlinfoRouter(Router *gin.RouterGroup) {
	HtmlinfoRouter := Router.Group("htmlinfo").Use(middleware.OperationRecord())
	{
		HtmlinfoRouter.POST("createHtmlinfo", v1.CreateHtmlinfo)   // 新建Htmlinfo
		HtmlinfoRouter.DELETE("deleteHtmlinfo", v1.DeleteHtmlinfo) // 删除Htmlinfo
		HtmlinfoRouter.DELETE("deleteHtmlinfoByIds", v1.DeleteHtmlinfoByIds) // 批量删除Htmlinfo
		HtmlinfoRouter.PUT("updateHtmlinfo", v1.UpdateHtmlinfo)    // 更新Htmlinfo
		HtmlinfoRouter.GET("findHtmlinfo", v1.FindHtmlinfo)        // 根据ID获取Htmlinfo
		HtmlinfoRouter.GET("getHtmlinfoList", v1.GetHtmlinfoList)  // 获取Htmlinfo列表
	}
}
