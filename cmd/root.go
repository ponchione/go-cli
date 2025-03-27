package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
		examples and usage of using your application. For example:
		
		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to demonstrate Cobra usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("Root command executed. Use --help for subcommands.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	// Cobra automatically adds a --help flag
	// You can add global flags here using rootCmd.PersistentFlags()

	if err := rootCmd.Execute(); err != nil {
		// Cobra prints the error, so we just exit
		// fmt.Fprintf(os.Stderr, "Error: %v\n", err) // Optionally print again
		os.Exit(1)
	}
}

// init() is called by Go before main()
// Use init functions in command files to add them to the root command.
func init() {
	// Here you will define your flags and configuration settings.
	// Example global flag (available to all subcommands):
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.my-cli-suite.yaml)")

	// Example local flag (only for the root command):
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Helper function to add subcommands (called from init() in subcommand files)
func addCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}
