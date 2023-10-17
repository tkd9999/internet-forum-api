package repository

import (
	"github.com/junshintakeda/internet-forum/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IPostRepository interface {
	GetAllPosts(posts *[]models.Post, threadId uint) error
	GetPostByID(post *models.Post, postId uint) error
	GetPostByUserID(posts *[]models.Post, userId uint) error
	CreatePost(post *models.Post) error
	UpdatePost(post *models.Post, postId uint) error
	DeletePost(postId uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) IPostRepository {
	return &postRepository{db}
}

func (pr *postRepository) GetAllPosts(posts *[]models.Post, threadId uint) error {
	if err := pr.db.Joins("User").Where("thread_id=?", threadId).Order("created_at").Find(posts).Error; err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) GetPostByID(post *models.Post, postId uint) error {
	if err := pr.db.First(post, postId).Error; err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) GetPostByUserID(posts *[]models.Post, userId uint) error {
	if err := pr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(posts).Error; err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) CreatePost(post *models.Post) error {
	if err := pr.db.Create(post).Error; err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) UpdatePost(post *models.Post, postId uint) error {
	result := pr.db.Model(post).Clauses(clause.Returning{}).Where("id=?", postId).Update("content", post.Content)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (pr *postRepository) DeletePost(postId uint) error {
	result := pr.db.Where("id=?", postId).Delete(&models.Post{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
