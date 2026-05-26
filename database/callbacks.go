package database

import (
	"log"

	"gorm.io/gorm"
)

func RegisterCallbacks(db *gorm.DB) {
	log.Println("GORM hooks registered via model methods (AfterCreate, BeforeDelete)")
}