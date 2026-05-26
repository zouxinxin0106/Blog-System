package models

import "gorm.io/gorm"

type Comment struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Content string `gorm:"type:text;not null" json:"content"`
	UserID uint   `gorm:"not null;index" json:"user_id"`
	User   User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	PostID uint   `gorm:"not null;index" json:"post_id"`
	Post   Post   `gorm:"foreignKey:PostID" json:"post,omitempty"`
}

func (Comment) TableName() string {
	return "comments"
}

func (c *Comment) BeforeDelete(tx *gorm.DB) error {
	var post Post
	if err := tx.First(&post, c.PostID).Error; err != nil {
		return err
	}

	var count int64
	tx.Model(&Comment{}).Where("post_id = ? AND id != ?", c.PostID, c.ID).Count(&count)

	if count == 0 {
		post.CommentStatus = "无评论"
		return tx.Model(&Post{}).Where("id = ?", post.ID).Update("comment_status", post.CommentStatus).Error
	}
	return nil
}