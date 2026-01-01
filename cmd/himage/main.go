package main

import (
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "himage",
	Short: "Advanced Image Processing CLI",
	Long: `himage is a fast and flexible CLI tool for image processing.
It supports resizing, format conversion, filtering, and watermarking.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
