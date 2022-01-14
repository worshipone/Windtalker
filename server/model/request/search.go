package request

import "gin-vue-admin/model"

type SearchSearch struct{
    model.Search
    PageInfo
}