package repo

import (
	"time"
)

type Config struct {
	Driver         string
	Host           string
	Port           uint
	Database       string
	InstanceName   string
	User           string
	Password       string
	ConnectTimeout time.Duration
	ReadTimeout    string
	WriteTimeout   string
	DialTimeout    *time.Duration
	MaxLifetime    *time.Duration
	MaxIdleTime    *time.Duration
	MaxIdleConn    *int
	MaxOpenConn    *int
	SSLMode        bool
}
