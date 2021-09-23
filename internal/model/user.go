package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	LatestLoginAt time.Time `gorm:"column:latest_login_at;DEFAULT:NULL;"`
	LatestLoginIP string    `gorm:"column:latest_login_ip;"`
	OtpSecret     string    `gorm:"column:otpSecret;"`
	IsEnable      *bool     `gorm:"column:is_enable;INDEX;NOT NULL;DEFAULT:true"`
}

// QueryUser for repository where condition
type QueryUser struct {
	User         *User
	UnbindedOnly bool
	ForUpdate    bool

	UserNames []string
}

// Where for repository where condition
func (opts *QueryUser) Where(db *gorm.DB) *gorm.DB {
	db = db.Where(opts.User)
	//db = db.Scopes(opts.Base.WhereWith("address"))
	//db = db.Scopes(opts.Sorting.Sort)
	//db = db.Scopes(opts.Pagination.LimitAndOffset)

	if len(opts.UserNames) > 0 {
		db = db.Where("addr IN (?)", opts.UserNames)
	}
	return db
}

// Clause ...
func (opts *QueryUser) Clause() (exps []clause.Expression) {
	if opts.ForUpdate {
		exps = append(exps, clause.Locking{
			Strength: "UPDATE OF address",
		})
	}

	return exps
}

func (opts *QueryUser) Preload(db *gorm.DB) *gorm.DB {

	// 短解，可以考慮新增參數決定是否 Preload 這麼多東西
	// db = db.Preload("BindSection.Addrs")

	return db
}

// UpdateUserWhereOpts ...
type UpdateUserWhereOpts struct {
	User      *User
	ForUpdate bool
}

// Clause ...
func (opts *UpdateUserWhereOpts) Clause() (exps []clause.Expression) {
	if opts.ForUpdate {
		exps = append(exps, clause.Locking{
			Strength: "UPDATE",
		})
	}

	return exps
}

// Where ...
func (opts *UpdateUserWhereOpts) Where(db *gorm.DB) *gorm.DB {
	db = db.Where(opts.User)

	return db
}

