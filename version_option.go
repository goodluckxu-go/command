package command

import (
	"context"
	"github.com/gookit/color"
)

type versionOption struct {
}

func (v *versionOption) Explain() *ExplainOption {
	return &ExplainOption{
		Commands: []string{"-V", "--version"},
		Notes:    "版本号",
	}
}

func (v *versionOption) Handle(ctx context.Context, args ...string) {
	color.Println(versionMsg())
}
