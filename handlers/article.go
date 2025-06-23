package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dheepssupreme/project.git/database"
	"github.com/dheepssupreme/project.git/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetArticles - Mendapatkan daftar artikel dengan pagination
func GetArticles(c *gin.Context) {
	db := database.GetDB()
	var articles []models.Article

	// Pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// Filter
	published := c.Query("published")
	featured := c.Query("featured")
	authorID := c.Query("author_id")
	tag := c.Query("tag")

	query := db.Preload("Author").Preload("Tags")

	if published == "true" {
		query = query.Where("published = ?", true)
	}
	if featured == "true" {
		query = query.Where("featured = ?", true)
	}
	if authorID != "" {
		query = query.Where("author_id = ?", authorID)
	}
	if tag != "" {
		query = query.Joins("JOIN article_tags ON articles.id = article_tags.article_id").
			Joins("JOIN tags ON article_tags.tag_id = tags.id").
			Where("tags.slug = ?", tag)
	}

	// Execute query
	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil artikel"})
		return
	}

	// Count total
	var total int64
	countQuery := db.Model(&models.Article{})
	if published == "true" {
		countQuery = countQuery.Where("published = ?", true)
	}
	if featured == "true" {
		countQuery = countQuery.Where("featured = ?", true)
	}
	if authorID != "" {
		countQuery = countQuery.Where("author_id = ?", authorID)
	}
	countQuery.Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

// GetArticle - Mendapatkan artikel berdasarkan slug
func GetArticle(c *gin.Context) {
	db := database.GetDB()
	slug := c.Param("slug")

	var article models.Article
	if err := db.Preload("Author").Preload("Tags").Where("slug = ?", slug).First(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	// Update views count jika published
	if article.Published {
		db.Model(&article).Update("views_count", article.ViewsCount+1)
		
		// Track view
		userID, exists := c.Get("user_id")
		var userUUID *uuid.UUID
		if exists && userID.(string) != "" {
			parsed, err := uuid.Parse(userID.(string))
			if err == nil {
				userUUID = &parsed
			}
		}

		articleView := models.ArticleView{
			ArticleID: article.ID,
			UserID:    userUUID,
		}
		db.Create(&articleView)
	}

	c.JSON(http.StatusOK, gin.H{"article": article})
}

// CreateArticle - Membuat artikel baru
func CreateArticle(c *gin.Context) {
	db := database.GetDB()
	userID, _ := c.Get("user_id")

	var req struct {
		Title         string   `json:"title" binding:"required,min=5"`
		Slug          string   `json:"slug" binding:"required"`
		Subtitle      *string  `json:"subtitle"`
		Content       string   `json:"content" binding:"required,min=10"`
		Excerpt       *string  `json:"excerpt"`
		CoverImageURL *string  `json:"cover_image_url"`
		CoverImageAlt *string  `json:"cover_image_alt"`
		Published     bool     `json:"published"`
		Featured      bool     `json:"featured"`
		AllowComments bool     `json:"allow_comments"`
		Tags          []string `json:"tags"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate slug format
	if !isValidSlug(req.Slug) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format slug tidak valid"})
		return
	}

	// Check slug uniqueness
	var existingArticle models.Article
	if err := db.Where("slug = ?", req.Slug).First(&existingArticle).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Slug sudah digunakan"})
		return
	}

	authorUUID, _ := uuid.Parse(userID.(string))
	article := models.Article{
		AuthorID:      authorUUID,
		Title:         req.Title,
		Slug:          req.Slug,
		Subtitle:      req.Subtitle,
		Content:       req.Content,
		Excerpt:       req.Excerpt,
		CoverImageURL: req.CoverImageURL,
		CoverImageAlt: req.CoverImageAlt,
		Published:     req.Published,
		Featured:      req.Featured,
		AllowComments: req.AllowComments,
		WordCount:     countWords(req.Content),
		ReadingTime:   calculateReadingTime(req.Content),
	}

	if req.Published {
		now := time.Now()
		article.PublishedAt = &now
	}

	// Save artikel
	if err := db.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat artikel"})
		return
	}

	// Handle tags
	if len(req.Tags) > 0 {
		for _, tagSlug := range req.Tags {
			var tag models.Tag
			if err := db.Where("slug = ?", tagSlug).First(&tag).Error; err != nil {
				// Create tag if not exists
				tag = models.Tag{
					Name: strings.Title(strings.ReplaceAll(tagSlug, "-", " ")),
					Slug: tagSlug,
				}
				db.Create(&tag)
			}
			db.Model(&article).Association("Tags").Append(&tag)
		}
	}

	// Load relationships
	db.Preload("Author").Preload("Tags").First(&article, article.ID)

	c.JSON(http.StatusCreated, gin.H{"article": article})
}

// UpdateArticle - Update artikel
func UpdateArticle(c *gin.Context) {
	db := database.GetDB()
	userID, _ := c.Get("user_id")
	slug := c.Param("slug")

	var article models.Article
	if err := db.Where("slug = ?", slug).First(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	// Check authorization
	authorUUID, _ := uuid.Parse(userID.(string))
	if article.AuthorID != authorUUID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Tidak ada akses untuk mengubah artikel ini"})
		return
	}

	var req struct {
		Title         string   `json:"title" binding:"required,min=5"`
		Subtitle      *string  `json:"subtitle"`
		Content       string   `json:"content" binding:"required,min=10"`
		Excerpt       *string  `json:"excerpt"`
		CoverImageURL *string  `json:"cover_image_url"`
		CoverImageAlt *string  `json:"cover_image_alt"`
		Published     bool     `json:"published"`
		Featured      bool     `json:"featured"`
		AllowComments bool     `json:"allow_comments"`
		Tags          []string `json:"tags"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	article.Title = req.Title
	article.Subtitle = req.Subtitle
	article.Content = req.Content
	article.Excerpt = req.Excerpt
	article.CoverImageURL = req.CoverImageURL
	article.CoverImageAlt = req.CoverImageAlt
	article.Published = req.Published
	article.Featured = req.Featured
	article.AllowComments = req.AllowComments
	article.WordCount = countWords(req.Content)
	article.ReadingTime = calculateReadingTime(req.Content)

	if req.Published && article.PublishedAt == nil {
		now := time.Now()
		article.PublishedAt = &now
	}

	if err := db.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate artikel"})
		return
	}

	// Update tags
	db.Model(&article).Association("Tags").Clear()
	if len(req.Tags) > 0 {
		for _, tagSlug := range req.Tags {
			var tag models.Tag
			if err := db.Where("slug = ?", tagSlug).First(&tag).Error; err != nil {
				tag = models.Tag{
					Name: strings.Title(strings.ReplaceAll(tagSlug, "-", " ")),
					Slug: tagSlug,
				}
				db.Create(&tag)
			}
			db.Model(&article).Association("Tags").Append(&tag)
		}
	}

	db.Preload("Author").Preload("Tags").First(&article, article.ID)
	c.JSON(http.StatusOK, gin.H{"article": article})
}

// DeleteArticle - Hapus artikel
func DeleteArticle(c *gin.Context) {
	db := database.GetDB()
	userID, _ := c.Get("user_id")
	slug := c.Param("slug")

	var article models.Article
	if err := db.Where("slug = ?", slug).First(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	// Check authorization
	authorUUID, _ := uuid.Parse(userID.(string))
	if article.AuthorID != authorUUID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Tidak ada akses untuk menghapus artikel ini"})
		return
	}

	if err := db.Delete(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus artikel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artikel berhasil dihapus"})
}

// Helper functions
func isValidSlug(slug string) bool {
	// Simple slug validation
	for _, char := range slug {
		if !((char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '-') {
			return false
		}
	}
	return len(slug) > 0
}

func countWords(content string) int {
	words := strings.Fields(content)
	return len(words)
}

func calculateReadingTime(content string) int {
	wordCount := countWords(content)
	// Assuming 200 words per minute reading speed
	readingTime := wordCount / 200
	if readingTime < 1 {
		readingTime = 1
	}
	return readingTime
} 