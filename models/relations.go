package models

import (
	"net"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Follow model untuk sistem following
type Follow struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	FollowerID  uuid.UUID `json:"follower_id" gorm:"type:uuid;not null"`
	FollowingID uuid.UUID `json:"following_id" gorm:"type:uuid;not null"`
	CreatedAt   time.Time `json:"created_at"`

	// Relationships
	Follower  Profile `json:"follower" gorm:"foreignKey:FollowerID"`
	Following Profile `json:"following" gorm:"foreignKey:FollowingID"`
}

// Like model untuk like artikel
type Like struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	ArticleID uuid.UUID `json:"article_id" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at"`

	// Relationships
	User    Profile `json:"user" gorm:"foreignKey:UserID"`
	Article Article `json:"article" gorm:"foreignKey:ArticleID"`
}

// Bookmark model untuk bookmark artikel
type Bookmark struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	ArticleID uuid.UUID `json:"article_id" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at"`

	// Relationships
	User    Profile `json:"user" gorm:"foreignKey:UserID"`
	Article Article `json:"article" gorm:"foreignKey:ArticleID"`
}

// CommentLike model untuk like komentar
type CommentLike struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	CommentID uuid.UUID `json:"comment_id" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at"`

	// Relationships
	User    Profile `json:"user" gorm:"foreignKey:UserID"`
	Comment Comment `json:"comment" gorm:"foreignKey:CommentID"`
}

// TagFollow model untuk follow tag
type TagFollow struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	TagID     uuid.UUID `json:"tag_id" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at"`

	// Relationships
	User Profile `json:"user" gorm:"foreignKey:UserID"`
	Tag  Tag     `json:"tag" gorm:"foreignKey:TagID"`
}

// ArticleView model untuk tracking views artikel
type ArticleView struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	ArticleID uuid.UUID  `json:"article_id" gorm:"type:uuid;not null"`
	UserID    *uuid.UUID `json:"user_id" gorm:"type:uuid"`
	IPAddress *net.IP    `json:"ip_address" gorm:"type:inet"`
	UserAgent *string    `json:"user_agent"`
	Referrer  *string    `json:"referrer"`
	Country   *string    `json:"country"`
	CreatedAt time.Time  `json:"created_at"`

	// Relationships
	Article Article  `json:"article" gorm:"foreignKey:ArticleID"`
	User    *Profile `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// Notification model untuk sistem notifikasi
type Notification struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	RecipientID uuid.UUID  `json:"recipient_id" gorm:"type:uuid;not null"`
	ActorID     *uuid.UUID `json:"actor_id" gorm:"type:uuid"`
	Type        string     `json:"type" gorm:"not null;check:type IN ('like', 'comment', 'follow', 'mention', 'article_published', 'welcome')"`
	Title       string     `json:"title" gorm:"not null"`
	Message     string     `json:"message" gorm:"not null"`
	ActionURL   *string    `json:"action_url"`
	ReadAt      *time.Time `json:"read_at"`
	CreatedAt   time.Time  `json:"created_at"`

	// Relationships
	Recipient Profile  `json:"recipient" gorm:"foreignKey:RecipientID"`
	Actor     *Profile `json:"actor,omitempty" gorm:"foreignKey:ActorID"`
}

// BeforeCreate hooks untuk generate UUID
func (f *Follow) BeforeCreate(tx *gorm.DB) error {
	if f.ID == uuid.Nil {
		f.ID = uuid.New()
	}
	return nil
}

func (l *Like) BeforeCreate(tx *gorm.DB) error {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return nil
}

func (b *Bookmark) BeforeCreate(tx *gorm.DB) error {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return nil
}

func (cl *CommentLike) BeforeCreate(tx *gorm.DB) error {
	if cl.ID == uuid.Nil {
		cl.ID = uuid.New()
	}
	return nil
}

func (tf *TagFollow) BeforeCreate(tx *gorm.DB) error {
	if tf.ID == uuid.Nil {
		tf.ID = uuid.New()
	}
	return nil
}

func (av *ArticleView) BeforeCreate(tx *gorm.DB) error {
	if av.ID == uuid.Nil {
		av.ID = uuid.New()
	}
	return nil
}

func (n *Notification) BeforeCreate(tx *gorm.DB) error {
	if n.ID == uuid.Nil {
		n.ID = uuid.New()
	}
	return nil
} 