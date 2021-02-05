package entity

import "gorm.io/gorm"

// Language table contains all the available languages
type Language struct {
	gorm.Model
	Name string `gorm:"unique;not null;size:2"`
	IETF string `gorm:"unique;not null;size:5"`
}

// Translation table contains all the transalation for the available languages
type Translation struct {
	gorm.Model
	LanguageName string   `gorm:"uniqueIndex:idx_translation"`
	Language     Language `gorm:"references:Name"`
	Scope        string   `gorm:"uniqueIndex:idx_translation;size:50"`
	Key          string   `gorm:"uniqueIndex:idx_translation;size:255"`
	Message      string   `gorm:"size:255;not null"`
}
