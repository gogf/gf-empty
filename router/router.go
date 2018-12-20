package router

import (
    "empty/app/controller/hello"
    "gitee.com/johng/gf/g"
)

// 统一路由注册.
func init() {
    g.Server().BindHandler("/", ctl_hello.Handler)
}
