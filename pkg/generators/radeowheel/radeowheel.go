package radeowheel

import (
	"fmt"
	"image"
	"math"
	"math/rand"

	"git.sr.ht/~sbinet/gg"
	"github.com/sovietscout/lapis/cmd/lapis"
)

var (
	radius = 4.0
	numOfDots = 4

	rings = 24

	colours = []string{
		"#cf9d9c",
		"#b9b2cd",
		"#9791d7",
		"#598adf",
		"#576cbf",
		"#4761a3",
		// "#213f81",
	}
)

type RadeoWheel lapis.Generator

func NewRadeoWheel(seed int64) *RadeoWheel {
	return &RadeoWheel{
		Name: "RadeoWheel",
		Seed: seed,
	}
}

func (g *RadeoWheel) Init() {
	g.Image = gg.NewContext(1000, 1000)
	rand.Seed(g.Seed)
}

func (g *RadeoWheel) Generate() image.Image {
	g.Image.SetHexColor("#0c172c")
	g.Image.Clear()

	g.Image.SetHexColor(getRandomColour(1))
	g.Image.DrawCircle(500, 500, radius)
	g.Image.Fill()

	for i := 1; i <= rings; i++ {

		for j := 0; j < numOfDots; j++ {
			rads := gg.Radians((360.0 / float64(numOfDots)) * float64(j))

			x := (math.Cos(rads) * radius * float64(i) * 4.5) + (1000 / 2)
			y := (math.Sin(rads) * radius * float64(i) * 4.5) + (1000 / 2)

			g.Image.SetHexColor(getRandomColour(i))
			g.Image.DrawCircle(x, y, radius)
			g.Image.Fill()
		}

		numOfDots += 5
	}

	return g.Image.Image()
}

func (g *RadeoWheel) GetFileName() string {
	return fmt.Sprintf("%s-%d.png", g.Name, g.Seed)
}

func getRandomColour(ringNo int) string {
	level := int(math.Ceil(float64(ringNo) / float64(len(colours))))
	index := level - 1

	coloursForRing := colours[index: index + 3]
	return coloursForRing[rand.Intn(len(coloursForRing))]
}