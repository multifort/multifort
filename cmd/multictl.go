package cmd

import (
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command{
	cmd := &cobra.Command{
		Use: 			   "multictl",
		Short:		       "multifort server framework",
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		Long:              `Start BeyondMesh API server`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	}

	cmd.AddCommand(newStartCmd())


	return cmd
}
