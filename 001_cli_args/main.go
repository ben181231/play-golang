package main

import (
	"flag"
	"fmt"
)

type argument struct {
	name     string
	language string
	years    int
}

func (a *argument) parse() {
	flag.StringVar(&a.name, "n", "", "Your name")
	flag.StringVar(&a.language, "l", "Golang", "The language you are learning")
	flag.IntVar(&a.years, "y", 1, "How many years do you learn the language")
	flag.Parse()
}

func main() {
	args := argument{}
	args.parse()

	fmt.Printf(
		"Hi %s! You have learnt %s for %d year(s).\n",
		args.name,
		args.language,
		args.years,
	)
}
