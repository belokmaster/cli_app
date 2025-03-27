package app

import "github.com/disintegration/imaging"

func ResizeImage(inputPath, outputPath string, width, height int) error {
	img, err := imaging.Open(inputPath)
	if err != nil {
		return err
	}

	resizedImg := imaging.Resize(img, width, height, imaging.Lanczos)
	return imaging.Save(resizedImg, outputPath)
}
