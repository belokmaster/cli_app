package main

import (
	"fmt"
	"os"

	"cli_app/internal/app"

	"github.com/spf13/cobra"
)

var (
	inputPath  string
	outputPath string
	width      int
	height     int
	angle      float64
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "imgtool",
		Short: "Image processing tool",
	}

	// команда обрезки
	cropCmd := &cobra.Command{
		Use:   "crop",
		Short: "Crop image",
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.CropImage(inputPath, outputPath, width, height); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Image cropped and saved to %s\n", outputPath)
		},
	}

	// команда поворота
	rotateCmd := &cobra.Command{
		Use:   "rotate",
		Short: "Rotate image",
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.RotateImage(inputPath, outputPath, angle); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Image rotated and saved to %s\n", outputPath)
		},
	}

	// команда чб
	grayscaleCmd := &cobra.Command{
		Use:   "grayscale",
		Short: "Convert image to grayscale",
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.GrayscaleImage(inputPath, outputPath); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Image converted to grayscale and saved to %s\n", outputPath)
		},
	}

	// команда изменения размера
	resizeCmd := &cobra.Command{
		Use:   "resize",
		Short: "Resize image",
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.ResizeImage(inputPath, outputPath, width, height); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Image resized and saved to %s\n", outputPath)
		},
	}

	// команда конвертации изображения
	convertCmd := &cobra.Command{
		Use:   "convert",
		Short: "Convert image format",
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.ConvertImage(inputPath, outputPath); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Image converted and saved to %s\n", outputPath)
		},
	}

	// общие флаги
	for _, cmd := range []*cobra.Command{cropCmd, rotateCmd, resizeCmd, convertCmd, grayscaleCmd} {
		cmd.Flags().StringVarP(&inputPath, "input", "i", "", "Input file path (required)")
		cmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output file path (required)")
		cmd.MarkFlagRequired("input")
		cmd.MarkFlagRequired("output")
	}

	// спец. флаги
	cropCmd.Flags().IntVarP(&width, "width", "W", 0, "Width (required)")
	cropCmd.Flags().IntVarP(&height, "height", "H", 0, "Height (required)")
	cropCmd.MarkFlagRequired("width")
	cropCmd.MarkFlagRequired("height")

	resizeCmd.Flags().IntVarP(&width, "width", "W", 0, "Width (0 for auto)")
	resizeCmd.Flags().IntVarP(&height, "height", "H", 0, "Height (0 for auto)")

	rotateCmd.Flags().Float64VarP(&angle, "angle", "a", 0, "Rotation angle in degrees")

	rootCmd.AddCommand(cropCmd, rotateCmd, resizeCmd, convertCmd, grayscaleCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
