package lapis

import (
	"image"
	"image/png"
	"os"
	"path/filepath"

	"git.sr.ht/~sbinet/gg"
)

type Generator struct {
	Name string
	Seed int64

	Image *gg.Context
}

type IGenerator interface {
	Init()
	Generate() image.Image
	GetFileName() string
}

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