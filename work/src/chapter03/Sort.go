package chapter03

import (
	"fmt"
	"sort"
)


/*
*   二分查找
*/
func SearchTest()  {
    x := 11
    s := []int{3, 6, 8, 11, 45} //注意已经升序排序
    pos := sort.Search(len(s), func(i int) bool { return s[i] >= x })
    if pos < len(s) && s[pos] == x {
        fmt.Println(x, "在s中的位置为：", pos)
    } else {
        fmt.Println("s不包含元素", x)
    }
}






func SortTest()  {
    stus := StuScores{
                {"alan", 95},
                {"hikerell", 91},
                {"acmfly", 96},
                {"leao", 90}}

    fmt.Println("Default:")
    //原始顺序
    for _, v := range stus {
        fmt.Println(v.name, ":", v.score)
    }
    fmt.Println()

    //StuScores已经实现了sort.Interface接口
    sort.Sort(stus) // 内部使用快速排序
    // sort.Sort(sort.Reverse(stus))    // Reverse其实就是把比较规则反过来
    
    fmt.Println("Sorted:")
     //排好序后的结构
    for _, v := range stus {
        fmt.Println(v.name, ":", v.score)
    }

    //判断是否已经排好顺序，将会打印true
    fmt.Println("IS Sorted?", sort.IsSorted(stus))
}

//学生成绩结构体
type StuScore struct {
     //姓名
    name  string
    //成绩
    score int
}

type StuScores []StuScore

//Len()
func (s StuScores) Len() int {
	return len(s)
}

//Less():成绩将有低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

//Swap()
func (s StuScores) Swap(i, j int) {
	 tem:=s[i]
     s[i] = s[j]
     s[j]=tem
}