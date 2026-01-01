package main

import (
	"github.com/ismailtsdln/himage/pkg/himage"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	filterType string
	filterVal  float64
)

var filterCmd = &cobra.Command{
	Use:   "filter [input-file/dir]",
	Short: "Apply a filter to the image",
	Long: `Apply various filters to the image.
Available filters: grayscale, blur, sharpen, invert, brightness, contrast, gamma, saturation, sepia, sigmoid`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ProcessInput(args[0], outputFile, func(inFile, outFile string) error {
			img, err := himage.Load(inFile)
			if err != nil {
				return err
			}

			switch filterType {
			case "grayscale":
				img.Grayscale()
			case "blur":
				img.Blur(filterVal)
			case "sharpen":
				img.Sharpen(filterVal)
			case "invert":
				img.Invert()
			case "brightness":
				img.AdjustBrightness(filterVal)
			case "contrast":
				img.AdjustContrast(filterVal)
			case "gamma":
				img.AdjustGamma(filterVal)
			case "saturation":
				img.AdjustSaturation(filterVal)
			case "sepia":
				img.Sepia()
			case "sigmoid":
				img.AdjustSigmoid(0.5, filterVal)
			default:
				// This will fail for each image in batch, maybe warn once?
				// For now, let it error out.
				pterm.Error.Println("Unknown filter type:", filterType)
				return nil // skip save
			}

			if outFile == "" {
				outFile = inFile
			}
			return img.SaveAs(outFile)
		})
	},
}

func init() {
	rootCmd.AddCommand(filterCmd)
	filterCmd.Flags().StringVarP(&filterType, "type", "t", "grayscale", "Filter type (grayscale, blur, brightness, etc.)")
	filterCmd.Flags().Float64VarP(&filterVal, "value", "v", 0, "Filter value (parameter for blur, brightness, etc.)")
	filterCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file path (or directory for batch)")
}
