package provisioning

import (
	"github.com/jonayrodriguez/translation-service/internal/translation/database/entity"
	"gorm.io/gorm"
)

func Languages(db *gorm.DB) {
	db.FirstOrCreate(&entity.Language{}, &entity.Language{Name: "uk", IETF: "en-UK"})
	db.FirstOrCreate(&entity.Language{}, &entity.Language{Name: "us", IETF: "en-US"})
}
