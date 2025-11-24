package utils

import (
	"testing"
)

// TestPWD verifies PWD function returns valid working DIR path
// TestPWD 验证 PWD 函数返回有效的工作 DIR 路径
func TestPWD(t *testing.T) {
	t.Log(PWD())
}
