package config

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"

	"github.com/yashmurty/wealth-park/wpark/pkg/logger"
	"github.com/yashmurty/wealth-park/wpark/pkg/util"
)

// Config contains all settings for the backend.
type Config struct {
	ServerID string `json:"server_id"` // Backend Server ID.
	Host     string `json:"host"`      // API server host.
	Port     int64  `json:"port"`      // API server port.

	IsProduction bool `json:"is_production"` // Indicates if the backend is running in production mode or not.

	MySQLURL    string `json:"mysql_url"`     // MySQL URL.
	MySQLDBName string `json:"mysql_db_name"` // MySQL DB Name.

	// Prints the config to stdout as pretty JSON.
	Debug bool
}

var (
	log           = logger.Get("config")
	config        *Config
	setupConfig   sync.Once
	whitespace, _ = regexp.Compile("[ \t,]+")
)

// GetInstance returns a singleton instance of Config.
func GetInstance() *Config {
	setupConfig.Do(func() {
		config = &Config{}
		config.loadFromEnv()
		config.DumpConfig()
	})
	return config
}

// DumpConfig prints the Config to stdout.
func (c *Config) DumpConfig() {
	if c.Debug {
		println("============ CONFIG DUMP ============:")
		println(util.GetPrettyJSON(c))
	}
}

// GetAddr returns a network address to the API gateway (format HOST:PORT).
func (c *Config) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// Load config from environment variables.
func (c *Config) loadFromEnv() {
	// Make sure the server id looks presentable.
	c.ServerID = GetEnv("WPARK_SERVER_ID", "wpark-server-1")
	c.Host = GetEnv("WPARK_CORE_API_HOST", "localhost")
	c.Port = GetEnvAsInt("WPARK_CORE_API_PORT", 11111)

	c.IsProduction = GetEnv("WPARK_PRODUCTION", "") != ""

	c.MySQLURL = GetEnv("WPARK_MYSQL_URL", "root:password@tcp(localhost:13306)")
	c.MySQLDBName = GetEnv("WPARK_MYSQL_DB_NAME", "wpark")

	c.Debug = GetEnv("WPARK_DUMP_CONFIG", "true") != "false"
}

// GetEnv looks up an environment variable.
func GetEnv(env, d string) string {
	if v, ok := os.LookupEnv(env); ok {
		return v
	}
	return d
}

// GetEnvAsInt looks up an environment variable and converts it to an int64.
func GetEnvAsInt(env string, d int64) int64 {
	n, err := strconv.ParseInt(os.Getenv(env), 0, 64)
	if err == nil {
		return n
	}
	return d
}
