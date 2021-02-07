package models

import (
	"novel/global"
)

type TbUser struct {
	Id         int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Username   string `gorm:"column:username;NOT NULL" json:"username"`
	Password   string `gorm:"column:password;NOT NULL" json:"password"`
	Integral   int    `gorm:"column:integral;default:0;NOT NULL" json:"integral"`
	CreateTime int    `gorm:"column:create_time;default:0;NOT NULL" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time;default:0;NOT NULL" json:"update_time"`
}

//绑定表名
func (t *TbUser) TableName() string {
	return "tb_user"
}

//添加用户
func (t *TbUser) Create() bool {
	err := orm.Create(t).Error
	if err != nil {
		global.Logger.Errorf("写入表TbUser失败：%v", err)
		return false
	}
	return true
}

func (t *TbUser) GetOne(id int, userName string, field string) *TbUser {
	command := orm.Select(field)
	if id > 0 {
		command = orm.Where("id=?", id)
	}
	if userName != "" {
		command = orm.Where("username=?", userName)
	}
	User := &TbUser{}
	command.Take(User)
	return User
}
