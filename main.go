package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	// Define the dimensions of the image
	width := 512
	height := 512

	// Create a new RGBA image with white background
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			// Calculate the distance from the current pixel to the center of the circle
			dx := x - width/2
			dy := y - height/2
			dist := dx*dx + dy*dy

			// Check if the distance is within the circle's radius
			if dist < height*height/10 {
				// Set the color of the pixel to red
				img.Set(x, y, color.RGBA{R: 255, G: 0, B: 0, A: 255})
			}
		}
	}

	// Save the image to a PNG file
	outputFile, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	if err := png.Encode(outputFile, img); err != nil {
		log.Fatal(err)
	}
}
