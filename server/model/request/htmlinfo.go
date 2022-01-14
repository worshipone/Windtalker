package request

import "gin-vue-admin/model"

type HtmlinfoSearch struct{
    model.Htmlinfo
    PageInfo
}