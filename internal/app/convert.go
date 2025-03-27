package app

import "github.com/disintegration/imaging"

func ConvertImage(inputPath, outputPath string) error {
	img, err := imaging.Open(inputPath)
	if err != nil {
		return err
	}

	return imaging.Save(img, outputPath)
}
