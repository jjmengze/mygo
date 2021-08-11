package gorm

import (
	"context"
	"gorm.io/gorm"
)

type RepoInterface interface {
	ReadDB(ctx context.Context) *gorm.DB
	WriteDB(ctx context.Context) *gorm.DB
}

type Repo struct {
	_readDB  *gorm.DB
	_writeDB *gorm.DB
}

func NewRepo(readDB, writeDB *gorm.DB) *Repo {
	return &Repo{
		_readDB:  readDB,
		_writeDB: writeDB,
	}
}
