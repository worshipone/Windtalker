package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitHtmltextLabeledRouter(Router *gin.RouterGroup) {
	HtmltextLabeledRouter := Router.Group("htmltextLabeled").Use(middleware.OperationRecord())
	{
		HtmltextLabeledRouter.POST("createHtmltextLabeled", v1.CreateHtmltextLabeled)   // 新建HtmltextLabeled
		HtmltextLabeledRouter.DELETE("deleteHtmltextLabeled", v1.DeleteHtmltextLabeled) // 删除HtmltextLabeled
		HtmltextLabeledRouter.DELETE("deleteHtmltextLabeledByIds", v1.DeleteHtmltextLabeledByIds) // 批量删除HtmltextLabeled
		HtmltextLabeledRouter.PUT("updateHtmltextLabeled", v1.UpdateHtmltextLabeled)    // 更新HtmltextLabeled
		HtmltextLabeledRouter.GET("findHtmltextLabeled", v1.FindHtmltextLabeled)        // 根据ID获取HtmltextLabeled
		HtmltextLabeledRouter.GET("getHtmltextLabeledList", v1.GetHtmltextLabeledList)  // 获取HtmltextLabeled列表
	}
}
