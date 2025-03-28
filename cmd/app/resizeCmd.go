package app

import (
	"cli_app/internal/image"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// команда изменения размера
var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Изменить размер изображения",
	Run: func(cmd *cobra.Command, args []string) {
		if err := image.ResizeImage(inputPath, outputPath, width, height); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Размер изображения изменён и сохранён в %s\n", outputPath)
	},
}

func init() {
	resizeCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Путь к исходному изображению (требуется)")
	resizeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Путь для сохранения нового изображения (требуется)")

	resizeCmd.MarkFlagRequired("input")
	resizeCmd.MarkFlagRequired("output")

	resizeCmd.Flags().IntVarP(&width, "width", "W", 0, "Ширина (0 = авто)")
	resizeCmd.Flags().IntVarP(&height, "height", "H", 0, "Высота (0 = авто)")
}
