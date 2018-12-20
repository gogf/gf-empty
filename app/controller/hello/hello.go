package ctl_hello

import (
    "gitee.com/johng/gf/g/net/ghttp"
)

// Hello World
func Handler(r *ghttp.Request) {
    r.Response.Writeln("Hello World!")
}
