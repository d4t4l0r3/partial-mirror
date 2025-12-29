package indexer

import (
	"context"
	"github.com/d4t4l0r3/partial-mirror/config"
	"github.com/d4t4l0r3/partial-mirror/ent"
	"github.com/charmbracelet/log"
	
	_ "github.com/lib/pq"
)

type DbConnector struct {
	client *ent.Client
}

func NewDbConnector(modifiers ...config.DbConnectorModifier) (DbConnector, error) {
	return NewDbConnectorWithModifiers(modifiers)
}

func NewDbConnectorWithModifiers(modifiers []config.DbConnectorModifier) (DbConnector, error) {
	conf := config.NewDbConnectorConfig()
	for _, modifier := range modifiers {
		modifier(&conf)
	}
	return NewDbConnectorWithConfig(conf)
}

func NewDbConnectorWithConfig(conf config.DbConnectorConfig) (connector DbConnector, err error) {
	log.Debug("Opening connection to DB", "host", conf.Host, "port", conf.Port, "user", conf.User, "dbname", conf.Database, "password", conf.Password, "sslmode", conf.SSLMode)
	connector.client, err = ent.Open("postgres", conf.ConnectionString())
	if err == nil {
		err = connector.client.Schema.Create(context.Background())
	}
	return
}

func (connector *DbConnector) Close() {
	connector.client.Close()
}
