# golang类似laravel的artisan的自定义命令生成工具

## 用法(usage)

引入
~~~go
import "github.com/goodluckxu-go/command"
~~~

命令文件(command)
~~~go
type Test struct {
}

func (t *Test) Explain() *command.ExplainCommand {
    return &command.ExplainCommand{
        Group:   "cache",
        Command: "clear",
        Notes:   "清除缓存",
        Help: func() {
            fmt.Println("清除缓存")
        },
}
}

func (t *Test) Handle(ctx context.Context, args ...string) {
    fmt.Println("清除缓存成功")
}
~~~

选项文件(option)
~~~go
type Kill struct {
}

func (k *Kill) Explain() *ExplainOption {
	return &ExplainOption{
		Commands: []string{"-k", "--kill"},
		Notes:    "删掉程序",
	}
}

func (k *Kill) Handle(ctx context.Context, args ...string) {
	fmt.Println("已经删除程序")
}
~~~

调用方法
~~~go
func main() {
	cmd := command.New()
	cmd.SetUsage("command [options] [arguments]") // 可设置，不设置默认值 command [options] [arguments]
    c.SetOptions(&Kill{}) // 设置选项，选项执行在命令之前
	cmd.SetCommands(&Test{}) // 设置命令
	cmd.Run(os.Args...)
}
~~~

展示效果
~~~shell
自定义命令 v1.00

Usage:
  command [options] [arguments]

Options:
  -V, --version  版本号
  -d, --daemon   守护进程执行
  -h, --help     帮助文档

Available commands:
 cache
  cache:clear  清除缓存
~~~