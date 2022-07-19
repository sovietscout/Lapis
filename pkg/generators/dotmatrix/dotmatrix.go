package dotmatrix

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"

	"git.sr.ht/~sbinet/gg"
	"github.com/sovietscout/lapis/cmd/lapis"
)

var (
	minDotSize = 0
	maxDotSize = 4
)

type DotMatrix lapis.Generator

func NewDotMatrix(seed int64) *DotMatrix {
	return &DotMatrix{
		Name: "DotMatrix",
		Seed: seed,
	}
}

func (g *DotMatrix) Init() {
	g.Image = gg.NewContext(1000, 1000)
	rand.Seed(g.Seed)
}

func (g *DotMatrix) Generate() image.Image {
	// Set background
	g.Image.SetHexColor("#333333")
	g.Image.Clear()

	// Draw dots
	for i := 0; i <= 1000; i += 100 {
		for j := 0; j <= 1000; j += 100 {
			g.Image.DrawImage(g.genDot(), i, j)
		}
	}

	return g.Image.Image()
}

func (g *DotMatrix) GetFileName() string {
	return fmt.Sprintf("%s-%d.png", g.Name, g.Seed)
}

func (g *DotMatrix) genDot() image.Image {
	// https://stackoverflow.com/a/54536595 (Generate rand num b/w range)
	radius := minDotSize + rand.Intn(maxDotSize - minDotSize + 1)

	dot := gg.NewContext(100, 100)
	dot.DrawCircle(50, 50, float64(radius * 3))
	dot.SetColor(color.White)
	dot.Fill()

	return dot.Image()
}