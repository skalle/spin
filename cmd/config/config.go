package config

import (
	"io"

	"github.com/spf13/cobra"
)

type configOptions struct {
}

var (
	configShort   = ""
	configLong    = ""
	configExample = ""
)

func NewConfigCmd(out io.Writer) *cobra.Command {
	options := configOptions{}
	cmd := &cobra.Command{
		Use:     "config",
		Aliases: []string{"config", "configuration", "co"},
		Short:   configShort,
		Long:    configLong,
		Example: configExample,
	}

	cmd.AddCommand(NewVerifyCmd(options))

	return cmd
}
