package main

import (
	"flag"
	"log"
	"os"

	"github.com/trk54ylmz/spritesheet/pkg/util"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(new(util.LogOutput))

	input := flag.String("input", "", "The input folder name is required.")
	output := flag.String("output", "", "The output file name is required.")

	flag.Parse()
	if flag.NFlag() == 0 || *input == "" || *output == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	err := Process(input, output)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("The file is created under \"%s\".", *output)
	os.Exit(0)
}
