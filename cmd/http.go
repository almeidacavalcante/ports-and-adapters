/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/almeidacavalcante/ports-and-adapters/adapters/web_server"
	"github.com/spf13/cobra"
)

var port string

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "CLI to start the HTTP server",
	Long: `Use this command to start the HTTP server. 
	Example: ports-and-adapters http`,
	Run: func(cmd *cobra.Command, args []string) {
		webServer := web_server.NewServer()
		webServer.ProductService = &productService
		println("Server started at http://localhost:9000")
		webServer.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	cliCmd.Flags().StringVarP(&port, "serve", "s", "9000", "Start the HTTP server")
}
