package lapis

import (
	"image"

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
