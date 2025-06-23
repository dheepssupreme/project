package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profile struct {
	ID                    uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Username              string    `json:"username" gorm:"uniqueIndex;not null"`
	FullName              *string   `json:"full_name"`
	Bio                   *string   `json:"bio"`
	AvatarURL             *string   `json:"avatar_url"`
	Website               *string   `json:"website"`
	Location              *string   `json:"location"`
	TwitterUsername       *string   `json:"twitter_username"`
	LinkedinUsername      *string   `json:"linkedin_username"`
	InstagramUsername     *string   `json:"instagram_username"`
	TiktokUsername        *string   `json:"tiktok_username"`
	FacebookUsername      *string   `json:"facebook_username"`
	FollowersCount        int       `json:"followers_count" gorm:"default:0"`
	FollowingCount        int       `json:"following_count" gorm:"default:0"`
	ArticlesCount         int       `json:"articles_count" gorm:"default:0"`
	TotalLikesReceived    int       `json:"total_likes_received" gorm:"default:0"`
	IsVerified            bool      `json:"is_verified" gorm:"default:false"`
	IsFeaturedWriter      bool      `json:"is_featured_writer" gorm:"default:false"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`

	// Relationships
	Articles      []Article      `json:"articles,omitempty" gorm:"foreignKey:AuthorID"`
	Followers     []Follow       `json:"followers,omitempty" gorm:"foreignKey:FollowingID"`
	Following     []Follow       `json:"following,omitempty" gorm:"foreignKey:FollowerID"`
	Likes         []Like         `json:"likes,omitempty" gorm:"foreignKey:UserID"`
	Bookmarks     []Bookmark     `json:"bookmarks,omitempty" gorm:"foreignKey:UserID"`
	Comments      []Comment      `json:"comments,omitempty" gorm:"foreignKey:AuthorID"`
	CommentLikes  []CommentLike  `json:"comment_likes,omitempty" gorm:"foreignKey:UserID"`
	TagFollows    []TagFollow    `json:"tag_follows,omitempty" gorm:"foreignKey:UserID"`
	Notifications []Notification `json:"notifications,omitempty" gorm:"foreignKey:RecipientID"`
}

func (p *Profile) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
} 