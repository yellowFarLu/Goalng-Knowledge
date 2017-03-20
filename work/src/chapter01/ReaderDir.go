package chapter01

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "strings"
    "os"
)

func ReadderDirTest() error {
    // fmt.Println(os.Args[1])
    if len(os.Args) > 1 {
		return Tree(os.Args[1], 1, map[int]bool{1:true})
	}

    return nil;
}

// 列出dirname目录中的目录树，实现类似Unix中的tree命令
// curHier 当前层级（dirname为第一层）
// hierMap 当前层级的上几层是否需要'|'的映射
func Tree(dirname string, curHier int, hierMap map[int]bool) error {
    // 获取该文件的绝对路径
    dirAbs, err := filepath.Abs(dirname)
    if err != nil {
        return err
    }

    // 获取文件夹的信息
    fileInfos, err := ioutil.ReadDir(dirAbs)
    if err != nil {
        return err
    }

    fileNum := len(fileInfos)
    
    
    for i, fileInfo := range fileInfos {
        for j := 1; j < curHier; j++ {
            if hierMap[j] {
                fmt.Print("|")
            } else {
                fmt.Print(" ")
            }
            fmt.Print(strings.Repeat(" ", 3))
        }
        
        // map是引用类型，所以新建一个map
        tmpMap := map[int]bool{}
        for k, v := range hierMap {
            tmpMap[k] = v
        }
        if i+1 == fileNum {
            fmt.Print("`")
            delete(tmpMap, curHier)
        } else {
            fmt.Print("|")
            tmpMap[curHier] = true
        }
        fmt.Print("-- ")
        fmt.Println(fileInfo.Name())
        if fileInfo.IsDir() {
            Tree(filepath.Join(dirAbs, fileInfo.Name()), curHier+1, tmpMap)
        }
    }
    return nil
}