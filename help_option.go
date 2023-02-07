package command

import (
	"context"
	"fmt"
	"github.com/gookit/color"
	"os"
	"os/exec"
	"path/filepath"
)

type helpOption struct {
}

func (h *helpOption) Explain() *ExplainOption {
	return &ExplainOption{
		Commands: []string{"-h", "--help"},
		Notes:    "帮助文档",
	}
}

func (h *helpOption) Handle(ctx context.Context, args ...string) {
	command := ""
	for _, arg := range args[1:] {
		if arg[0:1] != "-" {
			command = arg
			break
		}
	}
	if command == "" {
		s := "<fg=yellow>Usage:</>\n"
		s += "  只能存在一个Command，可以存在多个Option\n"
		s += "  <fg=red>开发时注意Option执行在Command之前</>\n"
		s += "<fg=yellow>Help:</>\n"
		s += "  <fg=green>-V, --version</>\n"
		s += "    查看命令版本\n"
		s += "  <fg=green>-d, --daemon</>\n"
		s += "    command -d 命来后台执行，目前只限于支持shell脚本的系统\n"
		s += "  <fg=green>-h, --help</>\n"
		s += "    [command] -h 查看命令帮助手册，命令定义在Help中\n"
		color.Print(s)
		os.Exit(0)
	}
	command += "/help"
	if os.Getppid() != 1 {
		//判断当其是否是子进程，当父进程return之后，子进程会被 系统1 号进程接管
		filePath, _ := filepath.Abs(args[0])   //将命令行参数中执行文件路径转换成可用路径
		cmd := exec.Command(filePath, command) //将其他命令传入生成出的进程
		output, err := cmd.CombinedOutput()
		if err != nil {
			errorMsg(err.Error())
		}
		fmt.Print(string(output))
	}
	os.Exit(0)
}
