package cmd

import (
	"database/sql"
	"github.com/almeidacavalcante/ports-and-adapters/adapters/db"
	"github.com/almeidacavalcante/ports-and-adapters/application"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"os"
)

func createProductService() application.ProductService {
	var database, err = sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		os.Exit(1)
	}
	var productDb = db.NewProductDb(database)
	var productService = application.ProductService{Persistence: productDb}
	return productService
}

var productService = createProductService()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ports-and-adapters",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ports-and-adapters.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
