package request

import "gin-vue-admin/model"

type DomaininfoSearch struct{
    model.Domaininfo
    PageInfo
}