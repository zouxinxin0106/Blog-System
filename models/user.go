package models

import "gorm.io/gorm"

type User struct {
	// Primitive fields like string, int, and time.Time are mapped to database columns
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Email     string    `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	PostCount uint      `gorm:"default:0" json:"post_count"`
	// Struct fields or slices of structs are mapped to associations in GORM
	// Association fields are not stored as table columns; GORM manages them through foreign keys
	// A struct is expanded into table columns only when using the gorm:"embedded" tag
	// Embedded fields are flattened into the parent table instead of creating associations
	Posts     []Post    `gorm:"foreignKey:UserID" json:"posts,omitempty"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) UpdatePostCount(db *gorm.DB) error {
	var count int64
	db.Model(&Post{}).Where("user_id = ?", u.ID).Count(&count)
	u.PostCount = uint(count)
	return db.Model(&User{}).Where("id = ?", u.ID).Update("post_count", u.PostCount).Error
}