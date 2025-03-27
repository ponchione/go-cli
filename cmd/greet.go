package cmd

import (
	"fmt"
	"go-cli/internal/greet"

	"github.com/spf13/cobra"
)

var (
	greetName   string
	greetFormal bool
	greetTimes  int
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",                                    // Command name: mycli greet
	Short: "Greets a specified person multiple times", // Short help description
	Long: `Prints a personalized greeting message to the console.

You can specify the name of the person to greet, how many times
the greeting should be repeated, and whether to use a formal tone.

Examples:
  mycli greet -n Alice                 # Simple greeting
  mycli greet --name Bob --times 3     # Greet Bob three times
  mycli greet -n Charlie -t 2 -f       # Greet Charlie formally twice`, // Long help description with examples
	Args: cobra.NoArgs,
	// Use RunE instead of Run to allow returning an error.
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1. Call the internal logic function with the flag values
		message, err := greet.GenerateGreeting(greetName, greetFormal, greetTimes)

		// 2. Handle potential errors from the logic function
		if err != nil {
			// Return the error so Cobra can display it
			return fmt.Errorf("failed to generate greeting: %w", err)
		}

		// 3. Print the successful result to standard output
		fmt.Println(message)

		// 4. Return nil to indicate success
		return nil
	},
}

// init() runs automatically when the package is initialized.
func init() {
	// Add the greetCmd to the rootCmd so it becomes available
	addCommand(greetCmd) // Using the helper function from root.go

	// Define the flags specific to the 'greet' command.
	// Flags are tied to the specific command (greetCmd here).
	greetCmd.Flags().StringVarP(&greetName, "name", "n", "", "Name of the person to greet (default: World)")
	greetCmd.Flags().BoolVarP(&greetFormal, "formal", "f", false, "Use a formal greeting ('Greetings' instead of 'Hello')")
	greetCmd.Flags().IntVarP(&greetTimes, "times", "t", 1, "Number of times to repeat the greeting")
}
