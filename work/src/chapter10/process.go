// 进程相关接口
package chapter10

import (
    "os/exec"
	"os"
	"fmt"
    "path/filepath"
	// "time"
	"log"
	"strings"
	"bytes"
	"io/ioutil"
	"os/user"
    // "syscall"
)

/*
创建线程
*/
func CreateThread()  {

}


/*
环境变量
*/
func Environ_Test()  {
    // 获取环境变量列表
    // fmt.Println(os.Environ())

    // 获取特定名称的环境变量
    fmt.Println(os.Getenv("PATH"))
}


/*
进程的当前工作目录
*/
func Dir_Test()  {
    // `Getwd` 返回一个对应当前工作目录的根路径。
    // 如果当前目录可以经过多条路径抵达（比如符号链接），`Getwd` 会返回其中一个。对应系统调用：`getcwd`
    fmt.Println(os.Getwd())
}



/*
包 `os/user` 允许通过名称或 ID 查询用户账号
*/
func UserTest()  {
    // 获取当前用户信息
    fmt.Println(user.Current())
    // 获取用户名对应的用户信息
    fmt.Println(user.Lookup("huangyuan"))
	// 获取用户Id对应的用户信息
    fmt.Println(user.LookupId("0"))
}


// os.Geteuid() 获取当前进程的有效用户 ID（effective user ID）
// os.Getegid() 有效组 ID（effectvie group ID）
func GetID_Test()  {
    id := os.Geteuid();
    gid := os.Getegid();
    fmt.Println(id, gid);
}




//因为 `Wait` 之后，会将管道关闭，所以，要使用这些管道相关的方法，只能使用 `Start`+`Wait` 组合，不能使用 `Run`。

/*
func (c *Cmd) StdoutPipe() (io.ReadCloser, error)       
StdoutPipe返回一个连接到command标准输出的管道pipe
*/
func StdoutPipeTest()  {
    cmd := exec.Command("ls")
    // 获取管道一定要在开始执行前
    stdout, _ := cmd.StdoutPipe()

    cmd.Start();

    // 读取管道一定要在输出之前
    content,_ := ioutil.ReadAll(stdout);
    fmt.Println(string(content))

    cmd.Wait()
}


/*
func (c *Cmd) StderrPipe() (io.ReadCloser, error)　　
StderrPipe返回一个pipe，这个管道连接到command的标准错误，当command命令退出时，Wait将关闭这些pipe

func (c *Cmd) StdinPipe() (io.WriteCloser, error)　　　
StdinPipe返回一个连接到command标准输入的管道pipe
*/
func StderrPipe_StdinPipe_Test()  {
    cmd := exec.Command("cat")
    // stdin, err := cmd.StdinPipe()
    // n, err:=stdin.Write([]byte("abc"))
    // cmd.Stdout = os.Stdout
    // stdin.Close()
    // fmt.Println(n,  err)

    // stderr,_ := cmd.StderrPipe();
    // var b []byte = make([]byte, 1024)
    // stderr.Read(b)
    // stderr.Close()
    // fmt.Println(b)

    cmd.Run();
}


// 注意：Output()和CombinedOutput()不能够同时使用，因为command的标准输出只能有一个，同时使用的话便会定义了两个，便会报错
/*
func (c *Cmd) Output() ([]byte, error)　　　　　//运行命令并返回其标准输出
*/
func OutputTest()  {
    cmd := exec.Command("ls")
    out, err := cmd.Output();
    fmt.Println(string(out), err)
}

/*
func (c *Cmd) CombinedOutput() ([]byte, error)　//运行命令，并返回标准输出和标准错误
*/
func CombinedOutputTest()  {
    cmd := exec.Command("ls");
    out,err := cmd.CombinedOutput();
    fmt.Println(string(out), err);
}


func FindProcess_Test()  {
    v,err := os.FindProcess(0)
    fmt.Println(v, err)
}

func Wait_Test()  {
    v,_ := os.FindProcess(128)
    ps,err := v.Wait();
    fmt.Println(ps, err)
}

/*运行外部程序, 使用StartProcess*/
func Test() {
    // filepath.Base() 返回路径的最后一个元素
    str := filepath.Base("/Users/huangyuan/Desktop/Golang/work/src/main");

    /* 在$PATH环境变量指定的路径中查找指定文件名（str）的可执行文件，然后返回这个文件的根路径
       试了很多次，只有可执行文件放在/usr/local/go/bin/目录下，才可以找到
    */
    lp, _ := exec.LookPath(str)
    // fmt.Println(lp, err);

    // ProcAttr保存新进程的属性
    procAttr := &os.ProcAttr{
        Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
    }
    // fmt.Println(procAttr)

    // 获取当前文件夹的一条根路径
    cwd,_ := os.Getwd();
    // fmt.Println(cwd, err)

    // 获取当前时间
    // start := time.Now()
    // fmt.Println(start)

    // 通过可执行文件的根路径，，预存的进程属性
    process,_ := os.StartProcess(lp, []string{cwd}, procAttr);
    // fmt.Println(process, err)

    processState, _ := process.Wait();
    fmt.Println(processState.String,  processState.SysUsage,  processState.UserTime, "哈哈")
}

func CommandTest()  {
    // name ：程序名字
    // args ：参数
    // 返回一个命令对象
    cmd := exec.Command("tr", "a-z", "A-Z")

    cmd.Stdin = strings.NewReader("my son");
    var out bytes.Buffer;
    cmd.Stdout = &out

    // 开始执行命令
    err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println(out)
}