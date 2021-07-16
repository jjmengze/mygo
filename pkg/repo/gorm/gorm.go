package gorm

import (
	"fmt"
	"gorm.io/gorm"
	"mygo/pkg/repo"
	"time"

	// database driver for gorm
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
)

// NewDB initialize a gorm.DB for further usage
func NewDatabase(config repo.Config) (*gorm.DB, error) {
	connectionFunc := supported[config.Driver]
	if connectionFunc == nil {
		return nil, fmt.Errorf("not supported DB driver : %s", config.Driver)
	}

	driver := connectionFunc(config)

	gormCfg := &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}
	engine, err := gorm.Open(driver, gormCfg)
	if err != nil {
		return nil, err
	}
	db, err := engine.DB()
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(*config.MaxLifetime)
	db.SetMaxOpenConns(*config.MaxOpenConn)
	db.SetConnMaxIdleTime(*config.MaxIdleTime)
	db.SetMaxIdleConns(*config.MaxIdleConn)
	return engine, nil
}

var supported = map[string]func(cfg repo.Config) gorm.Dialector{
	"cloudsql": func(cfg repo.Config) gorm.Dialector {
		return mysql.Open(fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=true&loc=UTC&time_zone=UTC&timeout=%s&readTimeout=%s&writeTimeout=%s",
			cfg.User, cfg.Password, cfg.InstanceName, cfg.Database, cfg.ConnectTimeout, cfg.ReadTimeout, cfg.WriteTimeout,
		))
	},
	"mysql": func(cfg repo.Config) gorm.Dialector {
		return mysql.Open(fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=UTC&time_zone=UTC&timeout=%s&readTimeout=%s&writeTimeout=%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.ConnectTimeout, cfg.ReadTimeout, cfg.WriteTimeout,
		))
	},
	"postgres": func(cfg repo.Config) gorm.Dialector {
		ssl := "disable"
		if cfg.SSLMode {
			ssl = "allow"
		}
		// TODO: 增加read,write,conn timeout
		return postgres.Open(fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s timezone=UTC",
			cfg.Host, cfg.Port, cfg.User, cfg.Database, cfg.Password, ssl,
		))
	},
	"sqlite": func(cfg repo.Config) gorm.Dialector {
		return sqlite.Open(fmt.Sprintf("%s.db", cfg.Database))
	},
}
