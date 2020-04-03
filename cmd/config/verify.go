package config

import (
	"github.com/spf13/cobra"
	"github.com/spinnaker/spin/cmd/gateclient"
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
	_, err := gateclient.NewGateClient(cmd.InheritedFlags())
	if err != nil {
		return err
	}
	return nil
}
