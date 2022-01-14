package request

import "gin-vue-admin/model"

type SettingSearch struct{
    model.Setting
    PageInfo
}