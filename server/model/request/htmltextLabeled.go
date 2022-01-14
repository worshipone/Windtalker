package request

import "gin-vue-admin/model"

type HtmltextLabeledSearch struct{
    model.HtmltextLabeled
    PageInfo
}