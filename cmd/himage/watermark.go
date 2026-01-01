package main

import (
	"image/color"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/streetbyters/himage/pkg/himage"
)

var (
	wmText    string
	wmImage   string
	wmX       int
	wmY       int
	wmOpacity float64
)

var watermarkCmd = &cobra.Command{
	Use:   "watermark [input-file/dir]",
	Short: "Add a watermark to the image",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ProcessInput(args[0], outputFile, func(inFile, outFile string) error {
			img, err := himage.Load(inFile)
			if err != nil {
				return err
			}

			if wmText != "" {
				// Default white text for now
				img.WatermarkText(wmText, wmX, wmY, color.White)
			} else if wmImage != "" {
				img.WatermarkImage(wmImage, wmX, wmY, wmOpacity)
				if img.Error() != nil {
					pterm.Error.Println("Failed to load watermark image:", img.Error())
					return img.Error()
				}
			} else {
				pterm.Error.Println("Please specify either --text or --image for the watermark")
				return nil
			}

			if outFile == "" {
				outFile = inFile
			}
			return img.SaveAs(outFile)
		})
	},
}

func init() {
	rootCmd.AddCommand(watermarkCmd)
	watermarkCmd.Flags().StringVarP(&wmText, "text", "T", "", "Watermark text")
	watermarkCmd.Flags().StringVarP(&wmImage, "image", "I", "", "Watermark image path")
	watermarkCmd.Flags().IntVarP(&wmX, "x", "x", 0, "X position")
	watermarkCmd.Flags().IntVarP(&wmY, "y", "y", 0, "Y position")
	watermarkCmd.Flags().Float64VarP(&wmOpacity, "opacity", "a", 0.5, "Opacity (0.0 - 1.0) for image watermark")
	watermarkCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file path (or directory for batch)")
}
