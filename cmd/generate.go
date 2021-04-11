package cmd

import (
	"github.com/spf13/cobra"

	"github.com/levibostian/dotenv/generate"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate lang file from .env values",
	Long: `Generate lang file from .env values. 

For example:

dotenv generate kotlin -p com.foo.foo`,
}

var packageName string
var outputFilePath string

// 1. Create a new command for your language
var kotlinGenerateCmd = &cobra.Command{
	Use:   "kotlin",
	Short: "Generate Env.kt",
	Run: func(cmd *cobra.Command, args []string) {
		generate.Execute("kotlin", outputFilePath)
	},
}

func init() {
	cobra.OnInitialize(generateInit)

	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&outputFilePath, "output", "o", "", "Output file path (example: build/) (optional. default: current dir)")

	// 2. Configure each language with their options
	kotlinGenerateCmd.Flags().StringVarP(&packageName, "packageName", "p", "", "Package name (example: com.foo.bar) (required)")
	kotlinGenerateCmd.MarkFlagRequired("packageName")

	// 3. Add commands for each langauge
	generateCmd.AddCommand(kotlinGenerateCmd)
}

func generateInit() {
	if outputFilePath == "" {
		outputFilePath = "./"
	}

	// TODO verify that outputFilePath is a real directory
}
