package config

import (
	"github.com/spf13/cobra"
	"github.com/spinnaker/spin/config"
)

type VerifyOptions struct {
	*configOptions
}

var (
	verifyConfigShort   = "Verify the configuration"
	verifyConfigLong    = "Verify the configuration"
	verifyConfigExample = "usage: spin config verify [options]"
)

func NewVerifyCmd(options configOptions) *cobra.Command {
	//options := VerifyOptions{
	//	configOptions: &options,
	//}
	cmd := &cobra.Command{
		Use:     "verify",
		Aliases: []string{"verify", "ve"},
		Short:   verifyConfigShort,
		Long:    verifyConfigLong,
		Example: verifyConfigExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			return verifyConfig(cmd, args)
		},
	}

	return cmd
}

func verifyConfig(cmd *cobra.Command, args []string) error {
	cfgPath, _, err := config.LoadConfig(cmd.InheritedFlags())
	if err != nil {
		return err
	}
	cmd.Println("ConfigPath:", cfgPath)
	cmd.Println("Config Load: success")
	return nil
}
