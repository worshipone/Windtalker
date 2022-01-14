package request

import "gin-vue-admin/model"

type LocationSearch struct{
    model.Location
    PageInfo
}