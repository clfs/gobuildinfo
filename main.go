package main

import (
	"debug/buildinfo"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0) // Disable output prefix.

	if len(os.Args) != 2 {
		log.Fatal("usage: gobuildinfo FILE")
	}

	info, err := buildinfo.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("failed to read build info: %v", err)
	}

	data, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		log.Fatalf("failed to convert build info into JSON: %v", err)
	}

	fmt.Printf("%s\n", data)
}
