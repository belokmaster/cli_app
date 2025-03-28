package app

import (
	"cli_app/internal/image"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Команда для конвертации (convert)
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Конвертировать изображение (PNG ↔ JPG)",
	Run: func(cmd *cobra.Command, args []string) {
		if err := image.ConvertImage(inputPath, outputPath); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Изображение конвертировано: %s\n", outputPath)
	},
}

func init() {
	convertCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Путь к исходному изображению (требуется)")
	convertCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Путь для сохранения нового изображения (требуется)")

	convertCmd.MarkFlagRequired("input")
	convertCmd.MarkFlagRequired("output")
}
