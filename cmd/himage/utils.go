package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
)

// ProcessInput handles both single file and directory inputs.
// It takes the input path, an output path (optional), and a processor function.
func ProcessInput(input string, output string, processor func(inFile, outFile string) error) {
	info, err := os.Stat(input)
	if err != nil {
		pterm.Error.Println("Failed to read input:", err)
		return
	}

	if info.IsDir() {
		// Batch processing
		spinner, _ := pterm.DefaultSpinner.Start("Batch processing directory...")

		err := filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			// Check extensions
			ext := strings.ToLower(filepath.Ext(path))
			if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" && ext != ".bmp" && ext != ".tiff" {
				return nil
			}

			// Determine output path for batch
			// If output is specified, mirror structure or dump there?
			// For simplicity, let's prefix "processed_" or use output dir if provided.
			// If output dir provided: output/relPath.

			var outPath string
			if output != "" {
				// assume output is a directory
				rel, _ := filepath.Rel(input, path)
				outPath = filepath.Join(output, rel)
				os.MkdirAll(filepath.Dir(outPath), 0755)
			} else {
				// In-place or prefix? Let's default to overwriting for now if explicit, but that's dangerous in batch.
				// Let's add a suffix.
				ext := filepath.Ext(path)
				outPath = strings.TrimSuffix(path, ext) + "_processed" + ext
			}

			if err := processor(path, outPath); err != nil {
				pterm.Warning.Println("Failed to process", path, ":", err)
			}
			return nil
		})

		if err != nil {
			spinner.Fail("Batch processing failed: " + err.Error())
		} else {
			spinner.Success("Batch processing completed")
		}

	} else {
		// Single file
		outPath := output
		if outPath == "" {
			outPath = input // Overwrite by default if not specified? Or safer to require output?
			// Existing commands defaulted to input or explicit output.
			// Let's keep it simple: passed straight to processor
			// But existing commands logic was slightly mixed.
			// Let's assume processor handles loading/saving logic, we just provide paths.
		}
		if err := processor(input, outPath); err != nil {
			pterm.Error.Println("Failed to process file:", err)
		}
	}
}
