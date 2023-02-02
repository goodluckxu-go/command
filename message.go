package command

import (
	"github.com/gookit/color"
	"os"
	"sort"
	"strings"
)

func help(c *Cmd) string {
	s := versionMsg()
	s += usageMsg(c.usage)
	s += optionMsg(c.options)
	return s
}

// 版本号
func versionMsg() string {
	return "自定义命令 <fg=green>v1.00</>\n\n"
}

// 用法
func usageMsg(opt string) string {
	s := ""
	if opt == "" {
		return s
	}
	s = "<fg=yellow>Usage:</>\n"
	s += "  " + opt + "\n\n"
	return s
}

// 选项
func optionMsg(opts []Option) string {
	s := ""
	if len(opts) == 0 {
		return s
	}
	s = "<fg=yellow>Options:</>\n"
	optMap := map[string]Option{}
	var cmdList []string
	for _, opt := range opts {
		cmd := optionKey(opt)
		optMap[cmd] = opt
		cmdList = append(cmdList, cmd)
	}
	sort.Strings(cmdList)
	max := maxLen(cmdList)
	for _, cmd := range cmdList {
		for _, opt := range opts {
			if cmd == optionKey(opt) {
				s += lineCommandNotes(max, 2, 2, optionKey(opt), opt.Explain().Notes) + "\n"
				break
			}
		}
	}
	s += "\n"
	return s
}

// 命令
func commandMsg(groupSearch string, opts []Command) string {
	s := ""
	if len(opts) == 0 {
		return s
	}
	s = "<fg=yellow>Available commands"
	if groupSearch != "" {
		s += " for the \"" + groupSearch + "\" namespace"
	}
	s += ":</>\n"
	optMap := map[string]Command{}
	var cmdList []string
	for _, opt := range opts {
		cmd := commandKey(opt)
		optMap[cmd] = opt
		cmdList = append(cmdList, cmd)
	}
	sort.Strings(cmdList)
	var noGroupList []Command
	var groupList []Command
	for _, cmd := range cmdList {
		for _, opt := range opts {
			if cmd == commandKey(opt) {
				if opt.Explain().Group == "" {
					noGroupList = append(noGroupList, opt)
				} else {
					groupList = append(groupList, opt)
				}
				break
			}
		}
	}
	max := maxLen(cmdList)
	for _, opt := range noGroupList {
		if groupSearch != "" {
			continue
		}
		s += lineCommandNotes(max, 2, 2, commandKey(opt), opt.Explain().Notes) + "\n"
	}
	group := ""
	for _, opt := range groupList {
		explan := opt.Explain()
		if groupSearch != "" && groupSearch != explan.Group {
			continue
		}
		if groupSearch == "" && explan.Group != group {
			group = explan.Group
			s += " <fg=yellow>" + group + "</>\n"
		}
		s += lineCommandNotes(max, 2, 2, commandKey(opt), opt.Explain().Notes) + "\n"
	}
	s += "\n"
	return s
}

func commandKey(opt Command) string {
	explan := opt.Explain()
	if explan.Group == "" {
		return explan.Command
	}
	return explan.Group + ":" + explan.Command
}

func optionKey(opt Option) string {
	explan := opt.Explain()
	var commands []string
	var oneLineList []string
	var twoLineList []string
	for _, c := range explan.Commands {
		removeLeftLine := strings.TrimLeft(c, "-")
		switch len(c) - len(removeLeftLine) {
		case 1:
			oneLineList = append(oneLineList, removeLeftLine)
		case 2:
			twoLineList = append(twoLineList, removeLeftLine)
		}
	}
	if len(oneLineList) > 0 {
		commands = append(commands, "-"+strings.Join(oneLineList, "|"))
	}
	if len(twoLineList) > 0 {
		commands = append(commands, "--"+strings.Join(twoLineList, "|"))
	}
	return strings.Join(commands, ", ")
}

func maxLen(opts []string) int {
	var max int
	for _, opt := range opts {
		if max < len(opt) {
			max = len(opt)
		}
	}
	return max
}

func lineCommandNotes(max, leftFill, rightFill int, cmd, notes string) string {
	middleFill := max - len(cmd) + rightFill
	s := fillData(leftFill, " ")
	s += "<fg=green>" + cmd + "</>"
	s += fillData(middleFill, " ")
	s += notes
	return s
}

func fillData(l int, sep string) string {
	s := ""
	for i := 0; i < l; i++ {
		s += sep
	}
	return s
}

func errorMsg(opts ...string) {
	max := maxLen(opts)
	s := "<bg=red>" + fillData(max+4, " ") + "</>\n"
	for _, opt := range opts {
		s += "<bg=red>" + fillData(2, " ") + "</>"
		s += "<fg=white;bg=red>" + opt + "</>"
		s += "<bg=red>" + fillData(2, " ") + "</>\n"
	}
	s += "<bg=red>" + fillData(max+4, " ") + "</>"
	color.Println(s)
	os.Exit(0)
}

func errorSuggest(cmd string, opts []Command) {
	group := strings.Split(cmd, ":")[0]
	s := "\n"
	s += "  <fg=white;bg=red> ERROR </> Command \"" + cmd + "\" is not defined."
	var suggestList []string
	for _, o := range opts {
		c := commandKey(o)
		if len(c)-len(strings.TrimPrefix(c, group)) > 0 {
			suggestList = append(suggestList, c)
		}
	}
	if len(suggestList) > 0 {
		s += " Did you mean one of these?\n"
		sort.Strings(suggestList)
	}
	s += "\n"
	for _, c := range suggestList {
		s += "  <fg=white>| " + c + "</>\n"
	}
	color.Println(s)
	os.Exit(0)
}
