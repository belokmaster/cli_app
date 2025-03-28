package app

import (
	"cli_app/internal/image"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// команда изменения размера
var cropCmd = &cobra.Command{
	Use:   "crop",
	Short: "Обрезать изображение",
	Run: func(cmd *cobra.Command, args []string) {
		if err := image.CropImage(inputPath, outputPath, width, height); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Изображение обрезано и сохранено в %s\n", outputPath)
	},
}

func init() {
	cropCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Путь к исходному изображению (требуется)")
	cropCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Путь для сохранения нового изображения (требуется)")

	cropCmd.MarkFlagRequired("input")
	cropCmd.MarkFlagRequired("output")

	cropCmd.Flags().IntVarP(&width, "width", "W", 0, "Ширина (0 = авто)")
	cropCmd.Flags().IntVarP(&height, "height", "H", 0, "Высота (0 = авто)")
}
