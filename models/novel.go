package models

import "novel/global"

type TbNovel struct {
	Id            int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Title         string `gorm:"column:title;NOT NULL" json:"title"`
	Image         string `gorm:"column:image;NOT NULL" json:"image"`
	ChapterSum    int    `gorm:"column:chapter_sum;default:0;NOT NULL" json:"chapter_sum"`
	ViewCount     int    `gorm:"column:view_count;default:0;NOT NULL" json:"view_count"`
	PublishStatus int    `gorm:"column:publish_status;default:0;NOT NULL" json:"publish_status"` // 发布状态
	Status        int    `gorm:"column:status;default:0;NOT NULL" json:"status"`                 // 连载状态
	CreateTime    int    `gorm:"column:create_time;default:0;NOT NULL" json:"create_time"`
	UpdateTime    int    `gorm:"column:update_time;default:0;NOT NULL" json:"update_time"`
}

//添加用户
func (t *TbNovel) Create() bool {
	err := orm.Create(t).Error
	if err != nil {
		global.Logger.Errorf("写入表TbNovel失败：%v", err)
		return false
	}
	return true
}
