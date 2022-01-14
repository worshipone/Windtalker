package request

import "gin-vue-admin/model"

type TaskManageSearch struct{
    model.TaskManage
    PageInfo
}