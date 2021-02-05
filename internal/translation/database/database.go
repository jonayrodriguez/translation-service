package database

import (
	"fmt"

	"github.com/jonayrodriguez/translation-service/internal/translation/config"
	"github.com/jonayrodriguez/translation-service/internal/translation/database/entity"
	"github.com/jonayrodriguez/translation-service/internal/translation/database/provisioning"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetInstance of the current database connection.
func GetInstance() *gorm.DB {
	return db
}

// InitDB to initialized the DB connection. Don´t required sync.Once.
func InitDB(c *config.DBConfig) error {
	connectionURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.Username, c.Password, c.Host, c.Port, c.Schema, c.Params)
	var err error
	// TODO- Custom the configuration (for example: Adding logging to the DB queries)
	db, err = gorm.Open(mysql.Open(connectionURL), &gorm.Config{})
	if err != nil {
		return err
	}

	// DDL auto migration
	err = db.AutoMigrate(&entity.Language{}, &entity.Translation{})
	if err != nil {
		return err
	}

	// Provisioning
	provisioning.Languages(db)

	return nil
}
