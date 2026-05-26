package models

import "gorm.io/gorm"

type Post struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Title          string    `gorm:"size:255;not null" json:"title"`
	Content        string    `gorm:"type:text" json:"content"`
	UserID         uint      `gorm:"not null;index" json:"user_id"`
	User           User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Comments       []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
	CommentStatus  string    `gorm:"size:50;default:'有评论'" json:"comment_status"`
}

func (Post) TableName() string {
	return "posts"
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	var user User
	if err := tx.First(&user, p.UserID).Error; err != nil {
		return err
	}
	return user.UpdatePostCount(tx)
}