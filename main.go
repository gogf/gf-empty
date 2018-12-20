package main

import (
	_ "empty/boot"
	_ "empty/router"
	"gitee.com/johng/gf/g"
)

func main() {
	g.Server().Run()
}
