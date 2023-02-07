package command

import "context"

type ExplainCommand struct {
	Group   string // 组
	Command string // 命令集合
	Notes   string // 注释
	Help    func() // 帮助
}

type Command interface {
	Explain() *ExplainCommand
	Handle(ctx context.Context, args ...string)
}

type ExplainOption struct {
	Commands []string // 命令集合
	Notes    string   // 注释
}

type Option interface {
	Explain() *ExplainOption
	Handle(ctx context.Context, args ...string)
}
