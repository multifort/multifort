package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command{
	cmd := &cobra.Command{
		Use: 			   "multictl",
		Short:		       "multifort server framework",
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		Long:              `Start multifort server`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	}

	cmd.AddCommand(newStartCmd())


	return cmd
}
