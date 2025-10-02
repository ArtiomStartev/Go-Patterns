package main

import (
	"fan-out-fan-in/image_processing"
	"image"
	"runtime"
	"strings"
	"sync"
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

	jobs := loadImages(imagePaths)

	// Fan-out this function to multiple goroutines
	resizedChan := resizeImages(jobs, 500, 500)

	// Collect / Fan-in results
	jobs = collectResults(resizedChan)

	saveImages(jobs)
}

func loadImages(paths []string) []Job {
	jobs := make([]Job, len(paths))

	for i, path := range paths {
		jobs[i] = Job{
			InputPath:  path,
			OutputPath: strings.Replace(path, "images/", "images/output/", 1),
			Image:      image_processing.ReadImage(path),
		}
	}

	return jobs
}

func resizeImages(jobs []Job, width, height int) <-chan Job {
	resizedChan := make(chan Job, len(jobs))
	limitChan := make(chan struct{}, runtime.NumCPU()) // Limit concurrency to number of CPU cores

	var wg sync.WaitGroup
	wg.Add(len(jobs))

	for _, job := range jobs {
		limitChan <- struct{}{} // Acquire a slot

		go func(job Job) {
			defer func() {
				<-limitChan // Release the slot
				wg.Done()
			}()

			job.Image = image_processing.ResizeImage(job.Image, width, height)
			resizedChan <- job
		}(job)
	}

	go func() {
		wg.Wait()
		close(resizedChan)
	}()

	return resizedChan
}

func collectResults(resizedChan <-chan Job) []Job {
	var jobs []Job
	for job := range resizedChan {
		jobs = append(jobs, job)
	}
	return jobs
}

func saveImages(jobs []Job) {
	for _, job := range jobs {
		image_processing.WriteImage(job.Image, job.OutputPath)
	}
}
