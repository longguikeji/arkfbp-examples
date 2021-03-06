package main

import (
	"fmt"
	"os"

	"github.com/longguikeji/arkfbp-go-examples/server"
	"github.com/spf13/cobra"
)

var (
	serverCmdParamHost string
	serveCmdParamPort  int
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the http server",
	Long:  `Start the http server.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(command *cobra.Command, args []string) {
		svr := server.NewHTTPServer()
		if err := svr.RegisterRoutes(Routes()); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err := svr.Serve(serverCmdParamHost, serveCmdParamPort); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	serveCmd.Flags().StringVarP(&serverCmdParamHost, "host", "", "127.0.0.1", "host to listen")
	serveCmd.Flags().IntVarP(&serveCmdParamPort, "port", "", 5000, "port to listen")

	RootCmd.AddCommand(serveCmd)
}
