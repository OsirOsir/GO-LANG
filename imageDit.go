package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/disintegration/imaging"
)

func main() {
	// Fetch badge image
	resp, err := http.Get("https://img.shields.io/badge/Go-46.3%25-blue")
	if err != nil {
		fmt.Println("Error fetching image:", err)
		return
	}
	defer resp.Body.Close()

	// Read image bytes
	imageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading image:", err)
		return
	}

	// Create a bytes.Buffer from imageBytes
	imageBuffer := bytes.NewBuffer(imageBytes)

	// Decode image
	srcImage, err := imaging.Decode(imageBuffer)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// Resize image to 300x100
	resizedImage := imaging.Resize(srcImage, 300, 100, imaging.Lanczos)

	// Save resized image
	err = imaging.Save(resizedImage, "resized_badge.png")
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}

	fmt.Println("Badge resized and saved successfully!")
}
