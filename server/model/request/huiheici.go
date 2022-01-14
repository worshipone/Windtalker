package request

import "gin-vue-admin/model"

type HuiheiciSearch struct{
    model.Huiheici
    PageInfo
}