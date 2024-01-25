package main

import (
	"fmt"
	"log"

	cfg "github.com/sebmaz93/TRG/pkg/config"
)

func main() {
	opts, err := cfg.GetOpts()
	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}

	fmt.Printf("opts: %+v", opts)
}
