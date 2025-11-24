// Package utils provides common functions used across wirekratos
// 包 utils 提供 wirekratos 中使用的通用函数
package utils

import (
	"strings"

	"github.com/yyle88/done"
	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexistpath/osmustexist"
)

// PWD gets current working DIR path via pwd command
// Returns absolute path with validation checks
//
// PWD 通过 pwd 命令获取当前工作 DIR 路径
// 返回经过验证检查的绝对路径
func PWD() string {
	// Execute pwd command and trim whitespace
	// 执行 pwd 命令并去除空白字符
	root := strings.TrimSpace(string(done.VAE(osexec.NewCMC().Exec("pwd")).Nice()))
	must.Nice(root)
	osmustexist.MustRoot(root)
	return root
}
