package command

import (
	"context"
	"github.com/gookit/color"
)

type version struct {
}

func (v *version) Explain() *ExplainOption {
	return &ExplainOption{
		Commands: []string{"-V", "--version"},
		Notes:    "版本号",
	}
}

func (v *version) Handle(ctx context.Context, args ...string) {
	color.Println(versionMsg())
}
