package request

import "gin-vue-admin/model"

type DemoSearch struct{
    model.Demo
    PageInfo
}