package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrDuplicateEmail = errors.New("邮箱冲突")
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type UserDAO struct {
	db *gorm.DB
}
type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 代表这是一个可以为 NULL 的列
	//Email    *string
	Email    string `gorm:"unique"`
	Password string
	// 创建时间
	Ctime int64
	// 更新时间
	Utime int64
}

func NewUserDao(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {

	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	err := dao.db.WithContext(ctx).Create(&u).Error

	if me, ok := err.(*mysql.MySQLError); ok {
		const duplicateErr = 1062
		if me.Number == duplicateErr {
			return ErrDuplicateEmail
		}
	}
	//fmt.Println("insert success")
	return err

}

func (dao *UserDAO) FindByEmail(ctx context.Context, email interface{}) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}
