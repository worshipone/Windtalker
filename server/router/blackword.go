package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitBlackwordRouter(Router *gin.RouterGroup) {
	BlackwordRouter := Router.Group("blackword").Use(middleware.OperationRecord())
	{
		BlackwordRouter.POST("createBlackword", v1.CreateBlackword)   // 新建Blackword
		BlackwordRouter.DELETE("deleteBlackword", v1.DeleteBlackword) // 删除Blackword
		BlackwordRouter.DELETE("deleteBlackwordByIds", v1.DeleteBlackwordByIds) // 批量删除Blackword
		BlackwordRouter.PUT("updateBlackword", v1.UpdateBlackword)    // 更新Blackword
		BlackwordRouter.GET("findBlackword", v1.FindBlackword)        // 根据ID获取Blackword
		BlackwordRouter.GET("getBlackwordList", v1.GetBlackwordList)  // 获取Blackword列表
	}
}
