package handlers

import (
	"blog-system/database"
	"blog-system/models"
)

func QueryTopPostsByCommentCount(limit int) ([]models.Post, error) {
	db := database.DB
	var posts []models.Post

	err := db.Model(&models.Post{}).
		// Use LEFT JOIN to include posts with zero comments
		// JOIN is used for query logic such as filtering, aggregation or sorting
		// comments is the table name, post_id is the foreign key in comments table, and posts.id is the primary key in posts table
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("COUNT(comments.id) DESC").
		// Preload is used for eager loading associations into Go structs
		Preload("User").
		// Comments is the association name defined in Post struct, and Comments.User is the nested association to load the user of each comment
		Preload("Comments").
		Preload("Comments.User").
		Limit(limit).
		Find(&posts).Error

	if err != nil {
		return nil, err
	}

	// return top N posts
	return posts, nil
}
