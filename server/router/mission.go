package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitMissionRouter(Router *gin.RouterGroup) {
	MissionRouter := Router.Group("mission").Use(middleware.OperationRecord())
	{
		MissionRouter.POST("createMission", v1.CreateMission)   // 新建Mission
		MissionRouter.DELETE("deleteMission", v1.DeleteMission) // 删除Mission
		MissionRouter.DELETE("deleteMissionByIds", v1.DeleteMissionByIds) // 批量删除Mission
		MissionRouter.PUT("updateMission", v1.UpdateMission)    // 更新Mission
		MissionRouter.GET("findMission", v1.FindMission)        // 根据ID获取Mission
		MissionRouter.GET("getMissionList", v1.GetMissionList)  // 获取Mission列表
	}
}
