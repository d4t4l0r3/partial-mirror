package main

import (
	"github.com/d4t4l0r3/partial-mirror/indexer"
	"github.com/charmbracelet/log"
)

func must[T any](obj T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return obj
}

func main() {
	log.SetLevel(log.DebugLevel)
	indexer := must(indexer.NewIndexer(indexer.WithSSLMode("disable")))
	indexer.Close()
}
