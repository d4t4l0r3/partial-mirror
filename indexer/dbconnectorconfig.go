package indexer

import "fmt"

type DbConnectorConfig struct {
	Host string
	Port int
	User string
	Database string
	Password string
	SSLMode string
}

func NewDbConnectorConfig() DbConnectorConfig {
	return DbConnectorConfig{"localhost", 5432, "mirror", "mirror", "dont-use-in-production", "require"}
}

func (config DbConnectorConfig) ConnectionString() string {
	return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v", config.Host, config.Port, config.User, config.Database, config.Password, config.SSLMode)
}

type DbConnectorModifier func(*DbConnectorConfig)

func WithHost(hostname string) DbConnectorModifier {
	return func(config *DbConnectorConfig) {
		config.Host = hostname
	}
}

func WithPort(port int) DbConnectorModifier {
	return func(config *DbConnectorConfig) {
		config.Port = port
	}
}

func WithUser(username string) DbConnectorModifier {
	return func(config *DbConnectorConfig) {
		config.User = username
	}
}

func WithDatabase(database string) DbConnectorModifier {
	return func(config *DbConnectorConfig) {
		config.Database = database
	}
}

func WithPassword(password string) DbConnectorModifier {
	return func(config *DbConnectorConfig) {
		config.Password = password
	}
}

func WithSSLMode(sslmode string) DbConnectorModifier {
	return func(config *DbConnectorConfig) {
		config.SSLMode = sslmode
	}
}
