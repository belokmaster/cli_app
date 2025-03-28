package app

import "github.com/spf13/cobra"

var (
	inputPath  string
	outputPath string
	width      int
	height     int
	angle      float64
)

var rootCmd = &cobra.Command{
	Use:   "imgtool",
	Short: "Утилита для обработки изображений",
}

func Execute() error {
	rootCmd.AddCommand(convertCmd)
	rootCmd.AddCommand(resizeCmd)
	rootCmd.AddCommand(rotateCmd)

	return rootCmd.Execute()
}
