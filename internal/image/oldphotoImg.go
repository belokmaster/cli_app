package image

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	"github.com/disintegration/imaging"
)

// MakeOldPhotoImage создаёт эффект старой фотографии
func MakeOldPhotoImage(inputPath, outputPath string) error {
	img, err := imaging.Open(inputPath)
	if err != nil {
		return err
	}

	// 1. Применяем сепию
	oldPhoto := applySepia(img)

	// 2. Уменьшаем насыщенность
	oldPhoto = imaging.AdjustSaturation(oldPhoto, -30)

	// 3. Добавляем шум (зернистость)
	oldPhoto = addNoise(oldPhoto, 0.0001)

	// 4. Уменьшаем контраст
	oldPhoto = imaging.AdjustContrast(oldPhoto, -10)

	// 5. Добавляем виньетирование
	oldPhoto = applyVignette(oldPhoto, 0.5)

	return imaging.Save(oldPhoto, outputPath)
}

// применяет эффект сепии к изображению
func applySepia(img image.Image) *image.NRGBA {
	bounds := img.Bounds()
	dst := imaging.New(bounds.Dx(), bounds.Dy(), color.NRGBA{0, 0, 0, 0})

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			newR := 0.393*float64(r>>8) + 0.769*float64(g>>8) + 0.189*float64(b>>8)
			newG := 0.349*float64(r>>8) + 0.686*float64(g>>8) + 0.168*float64(b>>8)
			newB := 0.272*float64(r>>8) + 0.534*float64(g>>8) + 0.131*float64(b>>8)

			// Ограничиваем значения
			newR = clampFloat(newR, 0, 255)
			newG = clampFloat(newG, 0, 255)
			newB = clampFloat(newB, 0, 255)

			dst.Set(x, y, color.NRGBA{
				R: uint8(newR),
				G: uint8(newG),
				B: uint8(newB),
				A: uint8(a >> 8),
			})
		}
	}

	return dst
}

// добавляет случайный шум
func addNoise(img image.Image, amount float64) *image.NRGBA {
	dst := imaging.Clone(img)
	bounds := dst.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if rand.Float64() < amount {
				r, g, b, a := dst.At(x, y).RGBA()
				noise := uint8(rand.NormFloat64() * 25)
				r = uint32(clamp(int(r>>8)+int(noise), 0, 255))
				g = uint32(clamp(int(g>>8)+int(noise), 0, 255))
				b = uint32(clamp(int(b>>8)+int(noise), 0, 255))
				dst.Set(x, y, color.NRGBA{
					R: uint8(r),
					G: uint8(g),
					B: uint8(b),
					A: uint8(a >> 8),
				})
			}
		}
	}

	return dst
}

// добавляет виньетирование
func applyVignette(img image.Image, strength float64) *image.NRGBA {
	bounds := img.Bounds()
	centerX := bounds.Dx() / 2
	centerY := bounds.Dy() / 2
	maxRadius := math.Sqrt(float64(centerX*centerX + centerY*centerY))

	dst := imaging.Clone(img)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			dist := math.Sqrt(float64((x-centerX)*(x-centerX) + (y-centerY)*(y-centerY)))
			darkness := 1 - strength*(dist/maxRadius)

			if darkness < 0 {
				darkness = 0
			}

			r, g, b, a := img.At(x, y).RGBA()
			r = uint32(float64(r>>8) * darkness)
			g = uint32(float64(g>>8) * darkness)
			b = uint32(float64(b>>8) * darkness)
			dst.Set(x, y, color.NRGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: uint8(a >> 8),
			})
		}
	}

	return dst
}

// ограничивает значение в пределах [min, max]
func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// ограничивает float
func clampFloat(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
