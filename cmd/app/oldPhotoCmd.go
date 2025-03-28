package app

import (
	"cli_app/internal/image"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Команда для конвертации (convert)
var oldphotoCmd = &cobra.Command{
	Use:   "oldphoto",
	Short: "Создание эффекта старой фотографии",
	Run: func(cmd *cobra.Command, args []string) {
		if err := image.MakeOldPhotoImage(inputPath, outputPath); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("На изображение наложен эффект старины: %s\n", outputPath)
	},
}

func init() {
	oldphotoCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Путь к исходному изображению (требуется)")
	oldphotoCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Путь для сохранения нового изображения (требуется)")

	oldphotoCmd.MarkFlagRequired("input")
	oldphotoCmd.MarkFlagRequired("output")
}
