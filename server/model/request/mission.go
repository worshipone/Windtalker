package request

import "gin-vue-admin/model"

type MissionSearch struct{
    model.Mission
    PageInfo
}