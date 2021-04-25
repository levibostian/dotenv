package cmd

import (
	cliConfig "github.com/levibostian/dotenv/cliconfig"
	"github.com/levibostian/dotenv/ui"
	"github.com/spf13/cobra"
)

var cfgFile string
var debug bool
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dotenv",
	Short: "Use .env values in your project.",
	Long:  `Use .env values in your project by generating a source code file you can compile inside your project. Works for multiple languages!`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ui.HandleError(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Show debug statements. Used for debugging program for bug reports and development. (default false)")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Show verbose logging. (default false)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	cliConfig.SetCliConfig(debug, verbose)
}
