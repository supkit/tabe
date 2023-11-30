package gorm

import (
	"fmt"
	"testing"
	"time"
)

type User struct {
	ID        uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT"` // 主键ID
	NickName  string    `gorm:"column:nickName;NOT NULL"`             // 昵称
	Gender    int       `gorm:"column:gender;default:0;NOT NULL"`     // 性别 1=男 2=女
	OpenId    string    `gorm:"column:openId;NOT NULL"`               // openId
	Source    string    `gorm:"column:source;NOT NULL"`               // 注册来源
	Avatar    string    `gorm:"column:avatar;NOT NULL"`               // 头像
	Email     string    `gorm:"column:email;NOT NULL"`                // 邮箱
	Mobile    string    `gorm:"column:mobile;NOT NULL"`               // 手机号
	DateBirth string    `gorm:"column:dateBirth;NOT NULL"`            // 出生年月
	Created   time.Time `gorm:"column:created"`                       // 创建时间
	Updated   time.Time `gorm:"column:updated"`                       // 更新时间
	Deleted   time.Time `gorm:"column:deleted"`                       // 删除时间
}

func (m *User) TableName() string {
	return "tb_user"
}

func TestNew(t *testing.T) {
	db, err := New("mysql.basic")
	fmt.Printf("DB %v err %v\n", db, err)

	var users []User
	err = db.Model(&users).Find(&users).Error

	fmt.Printf("user: %+v err: %v \n", users, err)
}
