package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	ArticleID    uuid.UUID  `json:"article_id" gorm:"type:uuid;not null"`
	AuthorID     uuid.UUID  `json:"author_id" gorm:"type:uuid;not null"`
	ParentID     *uuid.UUID `json:"parent_id" gorm:"type:uuid"`
	Content      string     `json:"content" gorm:"not null;check:length(trim(content)) >= 1"`
	LikesCount   int        `json:"likes_count" gorm:"default:0"`
	RepliesCount int        `json:"replies_count" gorm:"default:0"`
	IsEdited     bool       `json:"is_edited" gorm:"default:false"`
	IsPinned     bool       `json:"is_pinned" gorm:"default:false"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`

	// Relationships
	Article      Article       `json:"article" gorm:"foreignKey:ArticleID"`
	Author       Profile       `json:"author" gorm:"foreignKey:AuthorID"`
	Parent       *Comment      `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Replies      []Comment     `json:"replies,omitempty" gorm:"foreignKey:ParentID"`
	CommentLikes []CommentLike `json:"comment_likes,omitempty" gorm:"foreignKey:CommentID"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
} 