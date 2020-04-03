package config

import "github.com/spf13/cobra"

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
		Short:   verifyConfigShort,
		Long:    verifyConfigLong,
		Example: verifyConfigExample,
	}

	return cmd
}

func verifyConfig(cmd *cobra.Command, args []string) error {
	return nil
}
