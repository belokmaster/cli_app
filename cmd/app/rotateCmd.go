package app

import (
	"cli_app/internal/image"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// команда поворота
var rotateCmd = &cobra.Command{
	Use:   "rotate",
	Short: "Повернуть изображение",
	Run: func(cmd *cobra.Command, args []string) {
		if err := image.RotateImage(inputPath, outputPath, angle); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Image rotated and saved to %s\n", outputPath)
	},
}

func init() {
	rotateCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Путь к исходному изображению (требуется)")
	rotateCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Путь для сохранения нового изображения (требуется)")

	rotateCmd.MarkFlagRequired("input")
	rotateCmd.MarkFlagRequired("output")

	// Доп. флаг для поворота
	rotateCmd.Flags().Float64VarP(&angle, "angle", "a", 0, "Угол поворота (в градусах, например: 90)")
}
