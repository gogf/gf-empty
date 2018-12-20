package boot

import "gitee.com/johng/gf/g"

// 用于应用初始化。
func init() {
    g.Server().SetPort(8199)
}

