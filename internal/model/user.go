package model

import (
	"gorm.io/gorm"
	"time"
)

// User ...
type User struct {
	gorm.Model
	Password string  `gorm:"column:password;NOT NULL;"`
	Email    string  `gorm:"column:email;NOT NULL;"`
	Name     string  `gorm:"column:name;unique;type:varchar(100);not null;"`
	NikeName string  `gorm:"column:nikeName;unique;not null;index:;"`
	Age      uint8   `gorm:"column:age;"`
	Friends  []*User `gorm:"many2many:user_users;"` //foreignKey:name;References:name"` foreignKey//自己對到別人的外鍵，References別人身上的主鍵
	Height   uint8   `gorm:"column:height;"`
	Weight   uint8   `gorm:"column:weight;"`
	//posts         []Post
	LatestLoginAt time.Time `json:"latestLoginAt" gorm:"column:latest_login_at;DEFAULT:NULL;"`
	LatestLoginIP string    `json:"latestLoginIP" gorm:"column:latest_login_ip;"`
	IsEnable      *bool     `gorm:"column:is_enable;INDEX;NOT NULL;DEFAULT:true"`
}

// QueryUser for repository where condition
type QueryUser struct {
	Address      *User
	UnbindedOnly bool
	ForUpdate    bool

	NilActualCoinsOnly bool

	Addrs              []string
	CustomerServiceIDs []int64
	SectionIDs         []int64
}

// UpdateUserWhereOpts ...
type UpdateUserWhereOpts struct {
	Address   User
	ForUpdate bool
}
