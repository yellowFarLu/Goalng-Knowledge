package mysql
// 测试用例
import (
    "testing"
	"fmt"
)

func Test_findByPk(t *testing.T) {
    num := findByPk(1)
    // t.Log(num)
    fmt.Println(num)
}