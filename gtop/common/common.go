package common

import (
	"fmt"
	"github.com/bhmy-shm/tools/gtop/common/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const DefaultFormat = "gofks"

type Command struct {
	*cobra.Command
}

type Option func(*cobra.Command)

func WithRunE(runE func(*cobra.Command, []string) error) Option {
	return func(cmd *cobra.Command) {
		cmd.RunE = runE
	}
}

func WithRun(run func(*cobra.Command, []string)) Option {
	return func(cmd *cobra.Command) {
		cmd.Run = run
	}
}

func WithArgs(arg cobra.PositionalArgs) Option {
	return func(command *cobra.Command) {
		command.Args = arg
	}
}

// NewCommand 新增一个 cobra 命令
func NewCommand(use string, opts ...Option) *Command {
	c := &Command{
		Command: &cobra.Command{
			Use: use,
		},
	}
	for _, fn := range opts {
		fn(c.Command)
	}
	return c
}

func (c *Command) AddCommand(cmds ...*Command) {
	for _, cmd := range cmds {
		c.Command.AddCommand(cmd.Command)
	}
}

func (c *Command) Flags() *FlagSet {
	set := c.Command.Flags()
	return &FlagSet{
		FlagSet: set,
	}
}

// PersistentFlags 是cobra 库中 *cobra.Command 对象的一个方法，用于添加和获取那些应用于命令以及命令的所有子命令的标志（flags）。
/*
	例如，如果你有一个叫做 serve 的命令，它有一个子命令叫做 start，
	当你给 serve 命令添加一个持久标志时（比如 --port），这个 --port 标志也可以在 serve start 命令中使用。
*/
func (c *Command) PersistentFlags() *FlagSet {
	set := c.Command.PersistentFlags()
	return &FlagSet{
		FlagSet: set,
	}
}

// MustInit 用于初始化命令的帮助信息和用法
// 首先获取了命令本身及其所有递归的子命令，然后为每个命令设置 Short, Long, Example 文本，
// 以及遍历所有的标志（Flags 和 PersistentFlags）来设置它们的用法说明（Usage）
func (c *Command) MustInit() {
	commands := append([]*cobra.Command{c.Command}, getCommandsRecursively(c.Command)...)
	for _, command := range commands {
		commandKey := getCommandName(command)
		if len(command.Short) == 0 {
			command.Short = flags.Get(commandKey + ".short")
		}
		if len(command.Long) == 0 {
			command.Long = flags.Get(commandKey + ".long")
		}
		if len(command.Example) == 0 {
			command.Example = flags.Get(commandKey + ".example")
		}
		command.Flags().VisitAll(func(flag *pflag.Flag) {
			flag.Usage = flags.Get(fmt.Sprintf("%s.%s", commandKey, flag.Name))
		})
		command.PersistentFlags().VisitAll(func(flag *pflag.Flag) {
			flag.Usage = flags.Get(fmt.Sprintf("%s.%s", commandKey, flag.Name))
		})
	}
}

func getCommandName(cmd *cobra.Command) string {
	if cmd.HasParent() {
		return getCommandName(cmd.Parent()) + "." + cmd.Name()
	}
	return cmd.Name()
}

func getCommandsRecursively(parent *cobra.Command) []*cobra.Command {
	var commands []*cobra.Command
	for _, cmd := range parent.Commands() {
		commands = append(commands, cmd)
		commands = append(commands, getCommandsRecursively(cmd)...)
	}
	return commands
}

type FlagSet struct {
	*pflag.FlagSet
}

func (f *FlagSet) StringVar(p *string, name string) {
	f.StringVarWithDefaultValue(p, name, "")
}

func (f *FlagSet) StringVarWithDefaultValue(p *string, name string, value string) {
	f.FlagSet.StringVar(p, name, value, "")
}

func (f *FlagSet) StringVarP(p *string, name, shorthand string) {
	f.StringVarPWithDefaultValue(p, name, shorthand, "")
}

func (f *FlagSet) StringVarPWithDefaultValue(p *string, name, shorthand string, value string) {
	f.FlagSet.StringVarP(p, name, shorthand, value, "")
}

func (f *FlagSet) BoolVar(p *bool, name string) {
	f.BoolVarWithDefaultValue(p, name, false)
}

func (f *FlagSet) BoolVarWithDefaultValue(p *bool, name string, value bool) {
	f.FlagSet.BoolVar(p, name, value, "")
}

func (f *FlagSet) BoolVarP(p *bool, name, shorthand string) {
	f.BoolVarPWithDefaultValue(p, name, shorthand, false)
}

func (f *FlagSet) BoolVarPWithDefaultValue(p *bool, name, shorthand string, value bool) {
	f.FlagSet.BoolVarP(p, name, shorthand, value, "")
}

func (f *FlagSet) IntVar(p *int, name string) {
	f.IntVarWithDefaultValue(p, name, 0)
}

func (f *FlagSet) IntVarWithDefaultValue(p *int, name string, value int) {
	f.FlagSet.IntVar(p, name, value, "")
}

func (f *FlagSet) StringSliceVarP(p *[]string, name, shorthand string) {
	f.FlagSet.StringSliceVarP(p, name, shorthand, []string{}, "")
}

func (f *FlagSet) StringSliceVarPWithDefaultValue(p *[]string, name, shorthand string, value []string) {
	f.FlagSet.StringSliceVarP(p, name, shorthand, value, "")
}

func (f *FlagSet) StringSliceVar(p *[]string, name string) {
	f.StringSliceVarWithDefaultValue(p, name, []string{})
}

func (f *FlagSet) StringSliceVarWithDefaultValue(p *[]string, name string, value []string) {
	f.FlagSet.StringSliceVar(p, name, value, "")
}
