package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tag struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name           string    `json:"name" gorm:"uniqueIndex;not null;check:name ~ '^[a-zA-Z0-9\\s&+-]+$'"`
	Slug           string    `json:"slug" gorm:"uniqueIndex;not null;check:slug ~ '^[a-z0-9-]+$'"`
	Description    *string   `json:"description"`
	Color          string    `json:"color" gorm:"default:'#6366f1';check:color ~ '^#[0-9a-fA-F]{6}$'"`
	ArticlesCount  int       `json:"articles_count" gorm:"default:0"`
	FollowersCount int       `json:"followers_count" gorm:"default:0"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	// Relationships
	Articles   []Article   `json:"articles,omitempty" gorm:"many2many:article_tags"`
	TagFollows []TagFollow `json:"tag_follows,omitempty" gorm:"foreignKey:TagID"`
}

func (t *Tag) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
} 