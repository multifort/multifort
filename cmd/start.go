package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"multifort/log"
	"multifort/pkg/router"
	"net/http"
)

type startCmd struct {
	Host       string
	Port       int
	ConfigFile string
	Loglevel   string
}

func newStartCmd() *cobra.Command {

	startConfig := &startCmd{}

	cmd := &cobra.Command{
		Use:     "start",
		Short:   "Start multifort server",
		Example: "multictl start -c config/multifort.toml",
		RunE: func(cmd *cobra.Command, args []string) error {
			viper.SetConfigFile(startConfig.ConfigFile)
			viper.SetConfigType("toml")
			if err := viper.ReadInConfig(); err != nil {
				panic(fmt.Errorf("parse config file fail: %s", err))
			}
			log.InitLogger()
			engine := gin.Default()
			router.Init(engine)
			server := &http.Server{
				Addr:    fmt.Sprintf("%s:%d", startConfig.Host, startConfig.Port),
				Handler: engine,
			}

			//start HTTP server
			if err := server.ListenAndServe(); err != nil {
				if err == http.ErrServerClosed {
					log.Error("Multifort server closed under request\n")
				} else {
					log.Error("Multifort server closed with exception", err.Error())
				}
			}
			return nil
		},
	}

	flag := cmd.Flags()
	flag.StringVar(&startConfig.Host, "host", "0.0.0.0", "start multifort server with host")
	flag.IntVar(&startConfig.Port, "port", 8083, "override multifort port")
	flag.StringVar(&startConfig.ConfigFile, "config", "config/multifort.toml", "override multifort config file")
	flag.StringVar(&startConfig.Loglevel, "log-level", "info", "override multifort server log level")

	return cmd
}
