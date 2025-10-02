package main

import (
	"image"
	imageprocessing "pipeline/image_processing"
	"strings"
)

type Job struct {
	InputPath  string
	OutputPath string
	Image      image.Image
}

func main() {
	imagePaths := []string{
		"images/image1.png",
		"images/image2.png",
		"images/image3.png",
		"images/image4.png",
	}

	channel1 := loadImages(imagePaths)
	channel2 := resizeImages(channel1)
	channel3 := convertImagesToGrayscale(channel2)
	results := saveImages(channel3)

	for success := range results {
		if success {
			println("Image processed and saved successfully.")
		} else {
			println("Failed to process and save image.")
		}
	}
}

func loadImages(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		for _, path := range paths {
			img := imageprocessing.ReadImage(path)
			outputPath := strings.Replace(path, "images/", "images/output/", 1)
			out <- Job{InputPath: path, OutputPath: outputPath, Image: img}
		}
		close(out)
	}()
	return out
}

func resizeImages(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			job.Image = imageprocessing.ResizeImage(job.Image, 500, 500)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertImagesToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImages(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input {
			imageprocessing.WriteImage(job.OutputPath, job.Image)
			out <- true
		}
		close(out)
	}()
	return out
}
