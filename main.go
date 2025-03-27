package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
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
	var rootCmd = &cobra.Command{
		Use:   "imgtool",
		Short: "Утилита для обработки изображений",
	}

	// Команда для поворота (rotate)
	var rotateCmd = &cobra.Command{
		Use:   "rotate",
		Short: "Повернуть изображение",
		Run: func(cmd *cobra.Command, args []string) {
			img, err := imaging.Open(inputPath)
			if err != nil {
				fmt.Println("Ошибка открытия файла:", err)
				return
			}

			// Поворачиваем на указанный угол
			rotatedImg := imaging.Rotate(img, angle, nil)
			err = imaging.Save(rotatedImg, outputPath)
			if err != nil {
				fmt.Println("Ошибка сохранения:", err)
				return
			}
			fmt.Println("Изображение повёрнуто и сохранено в", outputPath)
		},
	}

	// Команда для изменения размера (resize)
	var resizeCmd = &cobra.Command{
		Use:   "resize",
		Short: "Изменить размер изображения",
		Run: func(cmd *cobra.Command, args []string) {
			img, err := imaging.Open(inputPath)
			if err != nil {
				fmt.Println("Ошибка открытия файла:", err)
				return
			}

			// Масштабируем (0 = сохранить пропорции)
			resizedImg := imaging.Resize(img, width, height, imaging.Lanczos)
			err = imaging.Save(resizedImg, outputPath)
			if err != nil {
				fmt.Println("Ошибка сохранения:", err)
				return
			}
			fmt.Println("Размер изображения изменён и сохранён в", outputPath)
		},
	}

	// Команда для конвертации (convert)
	var convertCmd = &cobra.Command{
		Use:   "convert",
		Short: "Конвертировать изображение (PNG ↔ JPG)",
		Run: func(cmd *cobra.Command, args []string) {
			img, err := imaging.Open(inputPath)
			if err != nil {
				fmt.Println("Ошибка открытия файла:", err)
				return
			}

			// Сохраняем в новом формате
			err = imaging.Save(img, outputPath)
			if err != nil {
				fmt.Println("Ошибка сохранения:", err)
				return
			}
			fmt.Printf("Изображение конвертировано: %s → %s\n", filepath.Ext(inputPath), filepath.Ext(outputPath))
		},
	}

	// Флаги для всех команд
	for _, cmd := range []*cobra.Command{cropCmd, rotateCmd, resizeCmd, convertCmd} {
		cmd.Flags().StringVarP(&inputPath, "input", "i", "", "Входной файл (обязательно)")
		cmd.Flags().StringVarP(&outputPath, "output", "o", "", "Выходной файл (обязательно)")
		cmd.MarkFlagRequired("input")
		cmd.MarkFlagRequired("output")
	}

	// Доп. флаги для crop/resize
	cropCmd.Flags().IntVarP(&width, "width", "W", 0, "Ширина (обязательно)")
	cropCmd.Flags().IntVarP(&height, "height", "H", 0, "Высота (обязательно)")
	cropCmd.MarkFlagRequired("width")
	cropCmd.MarkFlagRequired("height")

	resizeCmd.Flags().IntVarP(&width, "width", "W", 0, "Ширина (0 = авто)")
	resizeCmd.Flags().IntVarP(&height, "height", "H", 0, "Высота (0 = авто)")

	// Доп. флаги для rotate
	rotateCmd.Flags().Float64VarP(&angle, "angle", "a", 0, "Угол поворота (в градусах, например: 90)")

	// Добавляем команды
	rootCmd.AddCommand(cropCmd, rotateCmd, resizeCmd, convertCmd)

	// Запуск CLI
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
