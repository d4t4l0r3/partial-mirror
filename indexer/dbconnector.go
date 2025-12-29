package indexer

import (
	"context"
	"github.com/d4t4l0r3/partial-mirror/ent"
	"github.com/charmbracelet/log"
	
	_ "github.com/lib/pq"
)

type DbConnector struct {
	client *ent.Client
}

func NewDbConnector(modifiers ...DbConnectorModifier) (DbConnector, error) {
	return newDbConnectorInternal(modifiers)
}

func newDbConnectorInternal(modifiers []DbConnectorModifier) (connector DbConnector, err error) {
	config := NewDbConnectorConfig()
	for _, modifier := range modifiers {
		modifier(&config)
	}
	log.Debug("Opening connection to DB", "host", config.Host, "port", config.Port, "user", config.User, "dbname", config.Database, "password", config.Password, "sslmode", config.SSLMode)
	connector.client, err = ent.Open("postgres", config.ConnectionString())
	if err == nil {
		err = connector.client.Schema.Create(context.Background())
	}
	return
}

func (connector *DbConnector) Close() {
	connector.client.Close()
}
