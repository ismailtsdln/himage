package main

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/streetbyters/himage/pkg/himage"
)

var infoCmd = &cobra.Command{
	Use:   "info [input-file]",
	Short: "Get information about an image",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]

		spinner, _ := pterm.DefaultSpinner.Start("Reading image info...")
		img, err := himage.Load(inputFile)
		if err != nil {
			spinner.Fail("Failed to load image: " + err.Error())
			return
		}
		spinner.Success("Image loaded")

		// Get file info for size
		fileInfo, err := os.Stat(inputFile)
		size := "Unknown"
		if err == nil {
			size = fmt.Sprintf("%d bytes", fileInfo.Size())
		}

		bounds := img.Image.Bounds()

		data := [][]string{
			{"Property", "Value"},
			{"Path", img.Path},
			{"Format Convert Option", img.Ext}, // Ext is just file extension, real format requires decode config
			{"Width", fmt.Sprintf("%d px", bounds.Dx())},
			{"Height", fmt.Sprintf("%d px", bounds.Dy())},
			{"File Size", size},
		}

		pterm.DefaultTable.WithHasHeader().WithData(data).Render()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
