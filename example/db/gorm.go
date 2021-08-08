package main

import (
	"github.com/jjmengze/mygo/pkg/repo"
	"github.com/jjmengze/mygo/pkg/repo/gorm"
	"time"
)

func main() {

	gorm.NewDatabase(repo.Config{
		RetryTime: 10,
		Debug:     true,
		Driver:    repo.MySQL,
		Host:      "127.0.0.1",
		Port:      3307,
		Database:  "silkrode_annie_test_v2",
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

}
