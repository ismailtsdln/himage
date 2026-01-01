package main

import (
	"github.com/spf13/cobra"
	"github.com/streetbyters/himage/pkg/himage"
)

var (
	resizeWidth  int
	resizeHeight int
	outputFile   string
)

var resizeCmd = &cobra.Command{
	Use:   "resize [input-file/dir]",
	Short: "Resize an image or batch of images",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ProcessInput(args[0], outputFile, func(inFile, outFile string) error {
			img, err := himage.Load(inFile)
			if err != nil {
				return err
			}

			img.Resize(resizeWidth, resizeHeight)

			if outFile == "" {
				outFile = inFile
			}
			return img.SaveAs(outFile)
		})
	},
}

func init() {
	rootCmd.AddCommand(resizeCmd)
	resizeCmd.Flags().IntVarP(&resizeWidth, "width", "w", 0, "Width of the output image")
	resizeCmd.Flags().IntVarP(&resizeHeight, "height", "x", 0, "Height of the output image")
	resizeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file path (or directory for batch)")
}
