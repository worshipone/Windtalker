package request

import "gin-vue-admin/model"

type DomaincheckSearch struct{
    model.Domaincheck
    PageInfo
}