package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

type DriverType string

const (
	// MySQL is DB Driver
	MySQL DriverType = "mysql"
	// Postgres is DB Driver
	Postgres DriverType = "postgres"
	// CloudSql is DB Driver
	CloudSql DriverType = "cloudSql"
	// SqlLite is DB Driver
	SqlLite DriverType = "sqlite"
)

const (
	DefaultRetryTime       = 5
	DefaultConnMaxLifetime = 14400 * time.Second
	DefaultMaxOpenConns    = 5
	DefaultMaxIdleTime     = 1000 * time.Second
	DefaultMaxIdleConn     = 2
)

type Config struct {
	RetryTime      int
	Debug          bool `yaml:"debug" mapstructure:"debug"`
	Driver         DriverType
	Host           string
	Port           uint
	Database       string
	InstanceName   string
	User           string
	Password       string
	SearchPath     string `yaml:"search_path" mapstructure:"search_path"` // pg should setting this value. It will restrict access to db schema.
	ConnectTimeout time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	DialTimeout    *time.Duration
	MaxLifetime    *time.Duration
	MaxIdleTime    *time.Duration
	MaxIdleConn    *int
	MaxOpenConn    *int
	SSLMode        bool
}

// GetConnection with Config Driver we can get connection string.
func GetConnection(config Config) (string, error) {
	var connection string
	switch config.Driver {
	case MySQL:
		connection = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=UTC&time_zone=UTC&timeout=%s&readTimeout=%s&writeTimeout=%s",
			config.User, config.Password, config.Host, config.Port, config.Database, config.ConnectTimeout, config.ReadTimeout, config.WriteTimeout,
		)
	case Postgres:
		connection = fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s timezone=UTC",
			config.Host, config.Port, config.User, config.Database, config.Password,
		)

		if config.SSLMode {
			connection += " sslmode=require"
		} else {
			connection += " sslmode=disable"
		}

		if strings.TrimSpace(config.SearchPath) != "" {
			connection = fmt.Sprintf("%s search_path=%s", connection, config.SearchPath)
		}
	case CloudSql:
		connection = fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=true&loc=UTC&time_zone=UTC&timeout=%s&readTimeout=%s&writeTimeout=%s",
			config.User, config.Password, config.InstanceName, config.Database, config.ConnectTimeout, config.ReadTimeout, config.WriteTimeout,
		)
	case SqlLite:
		connection = fmt.Sprintf("%s.db", config.Database)
	default:
		return "", errors.New("Not support driver")
	}
	return connection, nil
}

func SetOption(config Config, db *sql.DB) {
	if config.MaxLifetime != nil {
		db.SetConnMaxLifetime(*config.MaxLifetime)
	} else {
		db.SetConnMaxLifetime(DefaultConnMaxLifetime)
	}

	if config.MaxOpenConn != nil {
		db.SetMaxOpenConns(*config.MaxOpenConn)
	} else {
		db.SetMaxOpenConns(DefaultMaxOpenConns)
	}

	if config.MaxIdleTime != nil {
		db.SetConnMaxIdleTime(*config.MaxIdleTime)
	} else {
		db.SetConnMaxIdleTime(DefaultMaxIdleTime)
	}

	if config.MaxIdleConn != nil {
		db.SetMaxIdleConns(*config.MaxIdleConn)
	} else {
		db.SetMaxIdleConns(DefaultMaxIdleConn)
	}
}
