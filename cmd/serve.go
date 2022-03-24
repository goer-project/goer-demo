package cmd

import (
	"goer/bootstrap"
	"goer/global"

	"github.com/gin-gonic/gin"
	"github.com/goer-project/goer/server"
	"github.com/spf13/cobra"
)

var cmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start API server",
	Run:   runServe,
}

func init() {
	rootCmd.AddCommand(cmdServe)
}

func runServe(cmd *cobra.Command, args []string) {
	// Gin mode: debug, release, test
	if global.Config.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := bootstrap.SetupRouter()

	// http server
	server.Run(router, global.Config.App.Port)
}
