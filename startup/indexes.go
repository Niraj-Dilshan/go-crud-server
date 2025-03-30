package startup

import (
	"crud-server/api/task/model"
	"log"

	"gorm.io/gorm"
)

// InitIndexes sets up database indexes and automigration
func InitIndexes(db *gorm.DB) error {
	// Auto migrate models
	err := db.AutoMigrate(&model.Task{})
	if err != nil {
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}
