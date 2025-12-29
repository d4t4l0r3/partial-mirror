package config

import (
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
	"github.com/goccy/go-yaml"
)

func GetEnvOrDefault(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	} else {
		return defaultValue
	}
}

type Config struct {
	BindAddress string `yaml:"bind_address"`
	BindPort int `yaml:"bind_port"`
	ConfigFilePath string `yaml:"-"`
	UpstreamURLs []string `yaml:"upstreams"`
	SyncPackages []string `yaml:"packages"`
	RetentionCount int `yaml:"retention_count"`
	RetentionTime time.Duration `yaml:"-"`
	RetentionTimeString string `yaml:"retention_time"`
	DbConfig DbConnectorConfig `yaml:"db"`
}

func GetConfig() (config Config, err error) {
	config.BindAddress = GetEnvOrDefault("PARTIAL_MIRROR_BIND_ADDRESS", "localhost")
	
	portString := GetEnvOrDefault("PARTIAL_MIRROR_BIND_PORT", "80")
	config.BindPort, err = strconv.Atoi(portString)
	if err != nil {
		log.Error("Failed to parse port number", "PARTIAL_MIRROR_BIND_PORT", portString)
		return
	}

	config.ConfigFilePath = GetEnvOrDefault("PARTIAL_MIRROR_CONFIG_FILE", "/etc/partial-mirror/config.yaml")
	retentionCountString := GetEnvOrDefault("PARTIAL_MIRROR_RETENTION_COUNT", "0")
	config.RetentionCount, err = strconv.Atoi(retentionCountString)
	if err != nil {
		log.Error("Failed to parse retention count", "PARTIAL_MIRROR_RETENTION_COUNT", retentionCountString)
		return
	}

	retentionTimeString := GetEnvOrDefault("PARTIAL_MIRROR_RETENTION_TIME", "0")
	config.RetentionTime, err = time.ParseDuration(retentionTimeString)
	if err != nil {
		log.Error("Failed to parse retention time", "PARTIAL_MIRROR_RETENTION_TIME", retentionTimeString)
		return
	}

	config.DbConfig.Host = GetEnvOrDefault("PARTIAL_MIRROR_DB_HOST", "localhost")

	dbPortString := GetEnvOrDefault("PARTIAL_MIRROR_DB_PORT", "5432")
	config.DbConfig.Port, err = strconv.Atoi(dbPortString)
	if err != nil {
		log.Error("Failed to parse db port number", "PARTIAL_MIRROR_DB_PORT", dbPortString)
		return
	}

	config.DbConfig.User = GetEnvOrDefault("PARTIAL_MIRROR_DB_USER", "mirror")
	config.DbConfig.Database = GetEnvOrDefault("PARTIAL_MIRROR_DB_USER", "mirror")
	config.DbConfig.Password = GetEnvOrDefault("PARTIAL_MIRROR_DB_USER", "")
	config.DbConfig.SSLMode = GetEnvOrDefault("PARTIAL_MIRROR_DB_USER", "require")

	err = config.ReadFile()
	return
}

func (config *Config) ReadFile() (err error) {
	path := os.ExpandEnv(config.ConfigFilePath)
	log.Debug("Reading config file", "path", path)
	buffer, err := os.ReadFile(path)
	if err != nil {
		log.Error("Failed to open file", "path", config.ConfigFilePath)
		return
	}

	err = yaml.Unmarshal(buffer, config)
	if err != nil {
		log.Error("Failed to parse config", "config", string(buffer), "error", err)
	}

	config.RetentionTime, err = time.ParseDuration(config.RetentionTimeString)
	if err != nil {
		log.Error("Failed to parse retention time", "retention_time", config.RetentionTimeString)
	}

	return
}
