package handlers

import (
	"blog-system/database"
	"blog-system/models"
)

func QueryUserPosts(userID uint) ([]models.Post, error) {
	db := database.DB
	var posts []models.Post

	err := db.Preload("Comments").
		Preload("Comments.User").
		Where("user_id = ?", userID).
		Find(&posts).Error

	return posts, err
}