package image

import "github.com/disintegration/imaging"

func RotateImage(inputPath, outputPath string, angle float64) error {
	img, err := imaging.Open(inputPath)
	if err != nil {
		return err
	}

	rotatedImg := imaging.Rotate(img, angle, nil)
	return imaging.Save(rotatedImg, outputPath)
}
