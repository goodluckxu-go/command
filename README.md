# golang类似laravel的artisan的自定义命令生成工具

## 用法(usage)

引入
~~~go
import "github.com/goodluckxu-go/command"
~~~

命令文件
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

调用方法
~~~go
func main() {
	cmd := command.New()
	cmd.SetUsage("command [options] [arguments]") // 可设置，不设置默认值 command [options] [arguments]
	cmd.SetCommands(&Test{})
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