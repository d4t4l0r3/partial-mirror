package main

import (
	"github.com/d4t4l0r3/partial-mirror/config"
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
	conf := must(config.GetConfig())
	indx := must(indexer.IndexerFromConfig(conf))
	indx.Close()
}
