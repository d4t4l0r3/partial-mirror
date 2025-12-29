package indexer

import (
	"github.com/charmbracelet/log"
)

type Indexer struct {
	dbConnector DbConnector
}

func NewIndexer(modifiers ...DbConnectorModifier) (indexer Indexer, err error) {
	indexer.dbConnector, err = newDbConnectorInternal(modifiers)
	if err != nil {
		log.Error("Couldn't create DbConnector", "error", err)
	}
	return
}

func (indexer *Indexer) Close() {
	indexer.dbConnector.Close()
}
