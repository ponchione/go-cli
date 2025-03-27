package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-cli/internal/jupyterparser"
)

var (
	inputPath string
	//outputPath string
)

var jupyterParseCmd = &cobra.Command{
	Use:   "jp-parse",
	Short: "Parses one or all Jupyter Notebooks in a give directory",
	Long: `The [jp-parse] command will parse a Jupyter Notebook and extract all 
		   of the Python code. It will then create a new .py file with the same
		   name as the original Jupyter Notebook.  If the [inputPath] arg is a 
		   file, only that file will be parsed.  If the [inputPath] arg is a directory,
		   all Jupyter Notebooks found in the directory will be parsed.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		inputPath = args[0]
		err := jupyterparser.ProcessPath(inputPath)
		if err != nil {
			return fmt.Errorf("failed to parse Jupyter Notebook(s): %w", err)
		}

		fmt.Println("Jupyter Notebook parsing completed successfully.")

		return nil
	},
}

func init() {
	addCommand(jupyterParseCmd)

	jupyterParseCmd.Flags().StringVarP(&inputPath, "inputPath", "i", "", "Input path for Jupyter Notebooks")
	//jupyterParseCmd.Flags().StringVarP(&outputPath, "outputPath", "o", "", "Output path for generated files")
}
