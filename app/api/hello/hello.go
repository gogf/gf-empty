package hello

import (
    "github.com/gogf/gf/g/net/ghttp"
)

// Hello World
func Handler(r *ghttp.Request) {
    r.Response.Writeln("Hello World!")
}
