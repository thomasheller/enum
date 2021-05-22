package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func main() {
	var left, none bool
	var separator string

	flag.BoolVar(&left, "left", false, "align left instead of right (incompatible with -none)")
	flag.BoolVar(&none, "none", false, "align neither left nor right (incompatible with -left)")
	flag.StringVar(&separator, "separator", ". ", "separator")
	flag.Parse()

	if err := enumerate(bufio.NewScanner(os.Stdin), os.Stdout, left, none, separator); err != nil {
		log.Fatal(err)
	}
}
