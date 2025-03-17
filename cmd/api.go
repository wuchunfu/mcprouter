package cmd

import (
	"fmt"

	"github.com/chatmcp/mcprouter/router"
	"github.com/chatmcp/mcprouter/service/api"
	"github.com/chatmcp/mcprouter/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var apiConfigFile string

// startAPIServer starts the api server
func startAPIServer(port int) {
	s := api.NewAPIServer()

	s.Route(router.APIRoute)
	s.Start(port)
}

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "start api server",
	Long:  `start api server`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := util.InitConfigWithFile(apiConfigFile); err != nil {
			fmt.Printf("init config failed with file: %s, %v\n", apiConfigFile, err)
			return
		}

		port := viper.GetInt("api_server.port")
		if port == 0 {
			port = 8027
		}

		startAPIServer(port)
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.Flags().StringVarP(&apiConfigFile, "config", "c", ".env.toml", "config file (default is .env.toml)")
}
