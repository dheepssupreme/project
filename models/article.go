package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID               uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	AuthorID         uuid.UUID  `json:"author_id" gorm:"type:uuid;not null"`
	Title            string     `json:"title" gorm:"not null;check:length(title) >= 5"`
	Slug             string     `json:"slug" gorm:"uniqueIndex;not null;check:slug ~ '^[a-z0-9-]+$'"`
	Subtitle         *string    `json:"subtitle"`
	Content          string     `json:"content" gorm:"not null;check:length(content) >= 10"`
	Excerpt          *string    `json:"excerpt"`
	CoverImageURL    *string    `json:"cover_image_url"`
	CoverImageAlt    *string    `json:"cover_image_alt"`
	PublishedAt      *time.Time `json:"published_at"`
	ReadingTime      int        `json:"reading_time" gorm:"default:1;check:reading_time > 0"`
	WordCount        int        `json:"word_count" gorm:"default:0"`
	Published        bool       `json:"published" gorm:"default:false"`
	Featured         bool       `json:"featured" gorm:"default:false"`
	AllowComments    bool       `json:"allow_comments" gorm:"default:true"`
	ViewsCount       int        `json:"views_count" gorm:"default:0"`
	LikesCount       int        `json:"likes_count" gorm:"default:0"`
	CommentsCount    int        `json:"comments_count" gorm:"default:0"`
	BookmarksCount   int        `json:"bookmarks_count" gorm:"default:0"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`

	// Relationships
	Author       Profile       `json:"author" gorm:"foreignKey:AuthorID"`
	Tags         []Tag         `json:"tags,omitempty" gorm:"many2many:article_tags"`
	Likes        []Like        `json:"likes,omitempty" gorm:"foreignKey:ArticleID"`
	Bookmarks    []Bookmark    `json:"bookmarks,omitempty" gorm:"foreignKey:ArticleID"`
	Comments     []Comment     `json:"comments,omitempty" gorm:"foreignKey:ArticleID"`
	ArticleViews []ArticleView `json:"article_views,omitempty" gorm:"foreignKey:ArticleID"`
}

func (a *Article) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
} 