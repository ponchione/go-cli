package jupyterparser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Cell struct {
	CellType string   `json:"cell_type"`
	Source   []string `json:"source"`
}
type Notebook struct {
	Cells []Cell `json:"cells"`
}

func ParseJupyterNotebook(inputPath, outputPath string) error {

	notebookData, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("error reading notebook file: %v", err)
	}

	var notebook Notebook
	if err := json.Unmarshal(notebookData, &notebook); err != nil {
		return fmt.Errorf("error parsing notebook JSON: %v", err)
	}

	var codeBuilder strings.Builder
	codeBuilder.WriteString("# Code extracted from Jupyter Notebook" + filepath.Base(inputPath) + "\n")
	codeBuilder.WriteString("# Generated on: " + time.Now().Format("2006-01-02 15:04:05") + "\n\n")

	for i, cell := range notebook.Cells {
		if cell.CellType == "code" {
			codeBuilder.WriteString(fmt.Sprintf("# Cell %d (code)\n", i+1))
			cellCode := strings.Join(cell.Source, "")
			codeBuilder.WriteString(cellCode)

			if !strings.HasSuffix(cellCode, "\n") {
				codeBuilder.WriteString("\n")
			}
			codeBuilder.WriteString("\n")
		}
	}

	return nil
}

func ProcessNotebook(notebookPath string) error {
	baseName := filepath.Base(notebookPath)
	ext := filepath.Ext(baseName)
	nameWithoutExt := strings.TrimSuffix(baseName, ext)
	outputPath := filepath.Join(filepath.Dir(notebookPath), nameWithoutExt+".py")

	err := ParseJupyterNotebook(notebookPath, outputPath)
	if err != nil {
		return fmt.Errorf("error parsing Jupyter Notebook: %v", err)
	}
	return nil
}

// ProcessPath handles either a single file or a directory of notebooks
func ProcessPath(path string) error {
	// Get file info
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("error accessing path %s: %v", path, err)
	}

	// Check if it's a directory
	if fileInfo.IsDir() {
		// Walk through the directory
		processedCount := 0
		err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Skip directories and non-notebook files
			if info.IsDir() || filepath.Ext(filePath) != ".ipynb" {
				return nil
			}

			// Process this notebook
			if err := ProcessNotebook(filePath); err != nil {
				return err
			}

			processedCount++
			return nil
		})

		if err != nil {
			return fmt.Errorf("error walking directory %s: %v", path, err)
		}

		if processedCount == 0 {
			fmt.Println("No Jupyter notebooks found in the directory.")
		} else {
			fmt.Printf("Processed %d notebook(s) in directory %s\n", processedCount, path)
		}

		return nil
	} else {
		// Process single file
		if filepath.Ext(path) != ".ipynb" {
			return fmt.Errorf("input file %s is not a Jupyter notebook (.ipynb extension required)", path)
		}

		return ProcessNotebook(path)
	}
}
