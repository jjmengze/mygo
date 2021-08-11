package gorm

import (
	"errors"
	"github.com/jjmengze/mygo/pkg/backoffmanager"
	"github.com/jjmengze/mygo/pkg/repo"
	gormTelemetry "github.com/jjmengze/mygo/pkg/telemetry/gorm"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"k8s.io/klog"
	"time"

	// database driver for gorm
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
)

// NewDatabase initialize a gorm.DB for further usage
func NewDatabase(config repo.Config) (*gorm.DB, error) {
	backoff := backoffmanager.NewExponentialBackoffManager(time.Millisecond, time.Second, 10*time.Second, 1.1, 1.1, time.Now)
	connection, err := repo.GetConnection(config)
	if err != nil {
		return nil, err
	}
	driver, err := translationDialector(config.Driver, connection)
	if err != nil {
		return nil, err
	}

	gormCfg := &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}
	if config.Debug == true {
		gormCfg.Logger = logger.Default.LogMode(logger.Info)
	}

	var engine *gorm.DB
	retry := repo.DefaultRetryTime
	if config.RetryTime != 0 {
		retry = config.RetryTime
	}
	for i := 0; i < retry; i++ {
		<-backoff.Backoff().C
		engine, err = gorm.Open(driver, gormCfg)
		if err != nil {
			klog.Errorf(" %s open failed: %v", config.Driver, err)
		}
		db, err := engine.DB()
		if err != nil {
			klog.Errorf("%s get DB error: %v", config.Driver, err)
		}
		err = db.Ping()
		if err != nil {
			klog.Errorf("%s ping DB error: %v", config.Driver, err)
		}else {
			break
		}
	}

	if err != nil {
		return nil, err
	}
	db, err := engine.DB()
	if err != nil {
		return nil, err
	}
	repo.SetOption(config, db)

	return engine, nil
}

func translationDialector(driver repo.DriverType, connection string) (gorm.Dialector, error) {
	switch driver {
	case repo.MySQL:
		return mysql.Open(connection), nil
	case repo.Postgres:
		return postgres.Open(connection), nil
	case repo.CloudSql:
		return mysql.Open(connection), nil
	case repo.SqlLite:
		return sqlite.Open(connection), nil
	default:
		return nil, errors.New("Not support type")
	}
}

// GormWithPlugins Initialize gormTelemetry plugin with options
func GormWithPlugins(db *gorm.DB) error {
	plugin := gormTelemetry.NewPlugin()
	return db.Use(plugin)
}
