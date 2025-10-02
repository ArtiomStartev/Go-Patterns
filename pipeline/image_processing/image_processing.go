package imageprocessing

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func ReadImage(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	return img
}

func WriteImage(path string, img image.Image) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		panic(err)
	}
}

func Grayscale(img image.Image) image.Image {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(originalPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}
	return grayImg
}

func ResizeImage(img image.Image, width, height int) image.Image {
	return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
}
