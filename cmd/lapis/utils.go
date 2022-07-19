package lapis

import (
	"image"
	"image/png"
	"image/color"
	"os"
	"path/filepath"
)

func WriteToFile(filename string, img image.Image) error {
	outputDir := filepath.Join(".", "output")
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(outputDir, filename))
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

func AddMargin(img image.Image, x, y int, colour color.Color) image.Image {
	// TODO: Center img within margins with given size and colour
	return img
}
