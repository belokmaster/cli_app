package image

import (
	"image"

	"github.com/disintegration/imaging"
)

func CropImage(inputPath, outputPath string, width, height int) error {
	img, err := imaging.Open(inputPath)
	if err != nil {
		return err
	}

	croppedImg := imaging.Crop(img, image.Rect(0, 0, width, height))
	return imaging.Save(croppedImg, outputPath)
}
