package repo

import (
	"github.com/jjmengze/mygo/internal/model"
	"github.com/jjmengze/mygo/pkg/repo"
	infraGorm "github.com/jjmengze/mygo/pkg/repo/gorm"
	"gorm.io/gorm"
	"time"
)

// Repository ...
type Repository interface {
	UserRepository
}

type repository struct {
	_readDB  *gorm.DB
	_writeDB *gorm.DB
}

// NewRepository repository new constructor
func NewRepository(read, write *gorm.DB) Repository {
	return &repository{
		_readDB:  read,
		_writeDB: write,
	}
}

func NewGORM(config *repo.Config) (*gorm.DB, error) {
	db, err := infraGorm.NewDatabase(repo.Config{
		RetryTime: 10,
		Debug:     true,
		Driver:    repo.MySQL,
		Host:      "127.0.0.1",
		Port:      3306,
		Database:  "blog",
		//InstanceName:   "",//for cloud sql
		User:     "root",
		Password: "123456",
		//SearchPath:     "",//for pg

		ConnectTimeout: time.Second,      //todo fix to times
		ReadTimeout:    10 * time.Second, //todo fix to times
		WriteTimeout:   10 * time.Second, //todo fix to times

		//DialTimeout:    nil,//default setting
		//MaxLifetime:    nil,//default setting
		//MaxIdleTime:    nil,//default setting
		//MaxIdleConn:    nil,//default setting
		//MaxOpenConn:    nil,//default setting
		//SSLMode:        false, //for pg
	})
	db.AutoMigrate(&model.User{})
	return db, err
}
