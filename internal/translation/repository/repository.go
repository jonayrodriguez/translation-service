package repository

import (
	"strings"

	"github.com/jonayrodriguez/translation-service/internal/translation/database/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// TranslationRepository is the interface for the translation repository
type TranslationRepository interface {
	// GetTranslation to retrieve the translated message.
	GetTranslation(language, scope, keyPatter string) ([]entity.Translation, error)
	// AddTranslation to add a new translated message.
	AddTranslation(language, scope, key, translatedMessage string) (*entity.Translation, error)
	// UpdateTranslation to update an existing translated message.
	UpdateTranslation(language, scope, key, translatedMessage string) (*entity.Translation, error)
	// DeleteTranslation to delete an existing translated message.
	DeleteTranslation(language, scope, key string) error
}

// TranslationRepositoryImp struct used for the Translate repository implementation
type TranslationRepositoryImp struct {
	db *gorm.DB
}

// NewTranslateRepository creates new translate repository
func NewTranslateRepository(db *gorm.DB) (*TranslationRepositoryImp, error) {
	return &TranslationRepositoryImp{db: db}, nil
}

func (r *TranslationRepositoryImp) GetTranslation(language, scope, keyPattern string) ([]entity.Translation, error) {
	var translations []entity.Translation
	// This is protected against SQL injection
	if res := r.db.Find(&translations, "language_name = ? AND scope = ? AND REGEXP_LIKE(`key`,?)", language, scope, strings.Replace(keyPattern, "*", ".*", -1)); res.Error != nil {
		return nil, status.Error(codes.Internal, res.Error.Error())
	}
	return translations, nil
}

func (r *TranslationRepositoryImp) AddTranslation(language, scope, key, translatedMessage string) (*entity.Translation, error) {
	dbTranslation := &entity.Translation{LanguageName: language, Scope: scope, Key: key, Message: translatedMessage}
	if res := r.db.Create(dbTranslation); res.Error != nil {
		return nil, status.Error(codes.Internal, res.Error.Error())
	}
	return dbTranslation, nil
}

func (r *TranslationRepositoryImp) UpdateTranslation(language, scope, key, translatedMessage string) (*entity.Translation, error) {
	dbTranslation := &entity.Translation{LanguageName: language, Scope: scope, Key: key}
	if res := r.db.First(dbTranslation); res.Error != nil {
		return nil, status.Error(codes.NotFound, res.Error.Error())
	}
	if res := r.db.Model(dbTranslation).Update("Message", translatedMessage); res.Error != nil {
		return nil, status.Error(codes.Internal, res.Error.Error())
	}
	dbTranslation.Message = translatedMessage
	return dbTranslation, nil
}

func (r *TranslationRepositoryImp) DeleteTranslation(language, scope, key string) error {
	if res := r.db.Delete(&entity.Translation{LanguageName: language, Scope: scope, Key: key}); res.Error != nil {
		return status.Error(codes.Internal, res.Error.Error())
	}
	return nil
}
