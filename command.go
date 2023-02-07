package command

import (
	"context"
	"github.com/gookit/color"
	"os"
	"regexp"
)

type Cmd struct {
	usage            string
	options          []Option
	commands         []Command
	execMap          map[string]interface{}
	runCommand       Command
	commandOtherOpts []string
	runOptions       []Option
	searchGroup      string
}

func New() *Cmd {
	return new(Cmd)
}

// 设置用法说明
func (c *Cmd) SetUsage(usage string) {
	c.usage = usage
}

// 设置参数
func (c *Cmd) SetOptions(opts ...Option) {
	c.options = append(c.options, opts...)
}

// 设置命令
func (c *Cmd) SetCommands(opts ...Command) {
	c.commands = append(c.commands, opts...)
}

// 执行输出
func (c *Cmd) Run(opts ...string) {
	// 增加系统参数
	c.SetOptions(&version{}, &daemon{})
	c.valid(opts...)
	if len(opts) > 1 && c.searchGroup == "" && c.Exec(opts...) {
		return
	}
	s := help(c)
	s += commandMsg(c.searchGroup, c.commands)
	color.Println(s)
}

func (c *Cmd) valid(opts ...string) {
	if c.execMap == nil {
		c.execMap = map[string]interface{}{}
	}
	optReg := regexp.MustCompile(`^-{1,2}[a-zA-Z]+[a-zA-Z0-9_-]*$`)
	for _, opt := range c.options {
		explain := opt.Explain()
		for _, cmd := range explain.Commands {
			if !optReg.MatchString(cmd) {
				errorMsg("Option format is \"^-{1,2}[a-zA-Z]+[a-zA-Z0-9_-]*$\"")
			}
			if c.execMap[cmd] != nil {
				errorMsg("Option \"" + cmd + "\" repeated")
			}
			c.execMap[cmd] = opt
		}
	}
	cmdReg := regexp.MustCompile(`^[a-zA-Z]+:*[a-zA-Z]+[a-zA-Z0-9_-]*$`)
	for _, opt := range c.commands {
		explain := opt.Explain()
		cmd := ""
		if explain.Group != "" {
			cmd += explain.Group + ":"
		}
		cmd += explain.Command
		if !cmdReg.MatchString(cmd) {
			errorMsg("Command format is \"^[a-zA-Z]+[a-zA-Z0-9_-]*$\"")
		}
		if c.execMap[cmd] != nil {
			errorMsg("Command \"" + cmd + "\" repeated")
		}
		c.execMap[cmd] = opt
	}
	for _, opt := range opts[1:] {
		if opt[0:1] == "-" {
			if cmd, ok := c.execMap[opt].(Option); ok {
				c.runOptions = append(c.runOptions, cmd)
			} else {
				errorMsg("The \"" + opt + "\" option does not exist.")
			}
		} else {
			if c.runCommand == nil {
				if cmd, ok := c.execMap[opt].(Command); ok {
					c.runCommand = cmd
				} else {
					for _, command := range c.commands {
						if command.Explain().Group == opt {
							c.searchGroup = opt
							return
						}
					}
					errorSuggest(opt, c.commands)
				}
			} else {
				c.commandOtherOpts = append(c.commandOtherOpts, opt)
			}
		}
	}
}

func (c *Cmd) Exec(opts ...string) bool {
	if c.runCommand == nil && len(c.runOptions) == 0 {
		return true
	}
	// 参数
	for _, opt := range c.runOptions {
		opt.Handle(context.Background(), opts...)
		os.Exit(0)
	}
	// 命令
	if c.runCommand != nil {
		c.runCommand.Handle(context.Background(), c.commandOtherOpts...)
	}
	return true
}
