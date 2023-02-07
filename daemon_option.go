package command

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
)

type daemonOption struct {
}

func (d *daemonOption) Explain() *ExplainOption {
	return &ExplainOption{
		Commands: []string{"-d", "--daemon"},
		Notes:    "守护进程执行",
	}
}

func (d *daemonOption) Handle(ctx context.Context, args ...string) {
	var newArgs []string
	for _, arg := range args {
		if arg == "-d" || arg == "--daemon" {
			continue
		}
		newArgs = append(newArgs, arg)
	}
	if len(newArgs[1:]) == 0 {
		return
	}
	if os.Getppid() != 1 {
		//判断当其是否是子进程，当父进程return之后，子进程会被 系统1 号进程接管
		filePath, _ := filepath.Abs(newArgs[0])       //将命令行参数中执行文件路径转换成可用路径
		cmd := exec.Command(filePath, newArgs[1:]...) //将其他命令传入生成出的进程
		cmd.Stdin = os.Stdin                          //给新进程设置文件描述符，可以重定向到文件中
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start() //开始执行新进程，不等待新进程退出
		return
	}
}
