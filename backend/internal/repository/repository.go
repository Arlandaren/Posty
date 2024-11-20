package repository

import (
	models "service/internal/shared/storage/dto"

	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	log.Println("NewRepository")
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreatePost(post *models.Post) error {
	result := r.db.Create(post)
	return result.Error
}

func (r *Repository) GetPost(id int64) (*models.Post, error) {
	var post models.Post
	result := r.db.First(&post, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}

func (r *Repository) GetAllPosts(cursor int64, limit int64) ([]*models.Post, error) {
	var posts []*models.Post
	query := r.db.Order("id desc")
	if cursor != 0 {
		query = query.Where("id < ?", cursor)
	}
	if limit != 0 {
		query = query.Limit(int(limit))
	}
	result := query.Find(&posts)
	return posts, result.Error
}

func (r *Repository) CommentPost(comment *models.Comment) error {
	result := r.db.Create(comment)
	return result.Error
}

func (r *Repository) GetAllComments(postId int64, cursor int64, limit int64) ([]*models.Comment, error) {
	var comments []*models.Comment
	query := r.db.Where("post_id = ?", postId).Order("id desc")
	if cursor != 0 {
		query = query.Where("id < ?", cursor)
	}
	if limit != 0 {
		query = query.Limit(int(limit))
	}
	result := query.Find(&comments)
	return comments, result.Error
}
