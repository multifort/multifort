package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type startCmd struct {
	Host       string
	Port       int
	ConfigFile string
}

func newStartCmd() *cobra.Command {

	startConfig := &startCmd{}

	cmd := &cobra.Command{
		Use:     "start",
		Short:   "Start multifort server",
		Example: "multictl start -c config/database.toml",
		RunE: func(cmd *cobra.Command, args []string) error {
			viper.SetConfigFile(startConfig.ConfigFile)
			if err := viper.ReadInConfig(); err != nil {
				panic(fmt.Errorf("parse config file fail: %s", err))
			}
			return nil
		},
	}

	flag := cmd.Flags()
	flag.StringVar(&startConfig.Host, "host", "0.0.0.0", "start multifort server with host")
	flag.IntVar(&startConfig.Port, "port", 8083, "override multifort port")
	flag.StringVar(&startConfig.ConfigFile, "config", "config/database.toml", "override multifort config file")

	return cmd
}
