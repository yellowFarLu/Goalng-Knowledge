package chapter09

import (
	"testing"
	"strings"
)

// 普通测试
// 这个文件是测试 strings 包下面的 Index函数
// func Test_Index(t *testing.T)  {
//     const s,sep,want = "chicken", "ken", 3;
//     got := strings.Index(s, sep);
//     if got!=want {
//         t.Errorf("Index(%q,%q) = %v; want %v", s, sep, got, want)
//     }
// }


// 表驱动测试
func Test_Index(t *testing.T)  {
    var tests = []struct{
        s string
        sep string
        out int
    } {
        // 测试数据
        {"", "", 0},
        {"", "a", -1},
        {"fo", "foo", -1},
        {"foo", "foo", 0},
        {"oofofoofooo", "f", 2},
        // etc
    }

    for _, test:=range tests {
        actual:= strings.Index(test.s, test.sep);
        if actual!=test.out {
            t.Errorf("Index(%q,%q) = %v; want %v",
                         test.s, test.sep, actual, test.out)
        }
    }
}