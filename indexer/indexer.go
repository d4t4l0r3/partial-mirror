package indexer

import (
	"github.com/d4t4l0r3/partial-mirror/config"
	"github.com/charmbracelet/log"
)

type Indexer struct {
	dbConnector DbConnector
}

func NewIndexer(modifiers ...config.DbConnectorModifier) (indexer Indexer, err error) {
	indexer.dbConnector, err = NewDbConnectorWithModifiers(modifiers)
	if err != nil {
		log.Error("Couldn't create DbConnector", "error", err)
	}
	return
}

func IndexerFromConfig(conf config.Config) (indexer Indexer, err error) {
	indexer.dbConnector, err = NewDbConnectorWithConfig(conf.DbConfig)
	if err != nil {
		log.Error("Couldn't create DbConnector", "error", err)
	}
	return
}

func (indexer *Indexer) Close() {
	indexer.dbConnector.Close()
}
