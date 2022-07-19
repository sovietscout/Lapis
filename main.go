package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/sovietscout/lapis/cmd/lapis"
	"github.com/sovietscout/lapis/pkg/generators/dotmatrix"
	"github.com/sovietscout/lapis/pkg/generators/radeowheel"
)

func main() {
	// Set and parse flags
	flag.Usage = func() {
		fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
		flag.PrintDefaults()
	}

	seedFlag := flag.Int64("s", time.Now().UnixNano(), "Input a seed for the random function")
	genFlag := flag.String("g", "dotmatrix", "Choose a generator")
	flag.Parse()

	// Set generator
	var generator lapis.IGenerator

	switch strings.ToLower(*genFlag) {
	case "dotmatrix":
		generator = dotmatrix.NewDotMatrix(*seedFlag)
	case "radeowheel":
		generator = radeowheel.NewRadeoWheel(*seedFlag)
	default:
		log.Fatalf("Unknown generator \"%s\"", *genFlag)
	}

	// Initialise generator
	generator.Init()

	// Generate
	img := generator.Generate()
	fileName := generator.GetFileName()

	if err := lapis.WriteToFile(fileName, img); err != nil {
		log.Fatalf("There has been an error: %s", err)
	}

	fmt.Printf("Image \"%s\" has been generated\n", fileName)
}