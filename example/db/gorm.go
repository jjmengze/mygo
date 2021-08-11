package main

import (
	"context"
	"fmt"
	"github.com/jjmengze/mygo/pkg/repo"
	infraGorm "github.com/jjmengze/mygo/pkg/repo/gorm"
	"github.com/jjmengze/mygo/pkg/telemetry"
	"gorm.io/gorm"
	"k8s.io/klog"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := telemetry.Config{
		Name:     "GORM_Tracing",
		EndPoint: "http://127.0.0.1:14268/api/traces",
		Jaeger: &telemetry.Jaeger{
			Mode: telemetry.Collector,
		},
	}
	flushTracer, err := telemetry.NewTelemetryProvider(ctx, c)
	if err != nil {
		klog.Warning("tracing config error:", err)
	}
	defer flushTracer()

	db, err := infraGorm.NewDatabase(repo.Config{
		RetryTime: 10,
		Debug:     true,
		Driver:    repo.MySQL,
		Host:      "127.0.0.1",
		Port:      3306,
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
	if err != nil {
		panic(err)
	}
	err = infraGorm.GormWithPlugins(db)
	if err != nil {
		panic(err)
	}

	//db change to debug mode
	db = db.Debug()

	// Migrate the schema
	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic(err.Error())
	}

	doGormOperations(ctx, db)

}

func doGormOperations(ctx context.Context, db *gorm.DB) {
	p := &Product{Code: "D42", Price: 100}

	db = db.WithContext(ctx)

	// Create
	if tx := db.Create(p); tx.Error != nil {
		klog.Errorf(tx.Error.Error())
	}

	// Update
	p.Price = 200
	if tx := db.Updates(p); tx.Error != nil {
		klog.Errorf(tx.Error.Error())
	}

	// Read
	var product Product
	if tx := db.First(&product, 1); tx.Error != nil {
		klog.Errorf(tx.Error.Error())
	}

	if tx := db.First(&product, "code = ?", "D42"); tx.Error != nil {
		klog.Errorf(tx.Error.Error())
	}

	// Delete
	if tx := db.Delete(p); tx.Error != nil {
		klog.Errorf(tx.Error.Error())
	}

	// this select should fail due to invalid table
	db.Exec("SELECT * FROM not_found")
	fmt.Println("Done")
}
