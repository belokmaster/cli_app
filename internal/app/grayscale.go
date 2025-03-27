package app

import "github.com/disintegration/imaging"

func GrayscaleImage(inputPath, outputPath string) error {
	img, err := imaging.Open(inputPath)
	if err != nil {
		return err
	}

	grayScaleImg := imaging.Grayscale(img)
	return imaging.Save(grayScaleImg, outputPath)
}
