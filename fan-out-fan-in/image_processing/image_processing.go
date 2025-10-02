package image_processing

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func ReadImage(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening image:", err)
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		panic(err)
	}

	return img
}

func WriteImage(img image.Image, path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating image file:", err)
		panic(err)
	}
	defer file.Close()

	if err = png.Encode(file, img); err != nil {
		fmt.Println("Error encoding image to PNG:", err)
		panic(err)
	}
}

func ResizeImage(img image.Image, width, height int) image.Image {
	return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
}
