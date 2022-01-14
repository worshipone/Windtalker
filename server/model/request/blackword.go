package request

import "gin-vue-admin/model"

type BlackwordSearch struct{
    model.Blackword
    PageInfo
}