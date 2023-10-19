package usecase

import (
	"github.com/junshintakeda/internet-forum/models"
	"github.com/junshintakeda/internet-forum/repository"
	"github.com/junshintakeda/internet-forum/validator"
)

type IPostUsecase interface {
	GetAllPosts(threadId uint) ([]models.PostResponse, error)
	GetPostByID(postId uint) (models.PostResponse, error)
	GetPostByUserID(userId uint) ([]models.PostResponse, error)
	CreatePost(post models.Post) (models.PostResponse, error)
	UpdatePost(post models.Post, postId uint) (models.PostResponse, error)
	DeletePost(postId uint) error
}

type postUsecase struct {
	pr repository.IPostRepository
	pv validator.IPostValidator
}

func NewPostUsecase(pr repository.IPostRepository, pv validator.IPostValidator) IPostUsecase {
	return &postUsecase{pr, pv}
}

func (pu *postUsecase) GetAllPosts(threadId uint) ([]models.PostResponse, err) {
	posts := []models.Post{}
	if err := pu.pr.GetAllPosts(&posts, threadId); err != nil {
		return nil, err
	}
	resPosts := []models.PostResponse{}
	for _, v := range posts {
		p := models.PostResponse{
			ID:        v.ID,
			Content:   v.Content,
			Username:  v.User.Username,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resPosts = append(resPosts, p)
	}
	return resPosts, nil
}

func (pu *postUsecase) GetPostByID(postId uint) (models.PostResponse, error) {
	post := models.Post{}
	if err := pu.pr.GetPostByID(&post, postId); err != nil {
		return models.PostResponse{}, err
	}
	resPost := models.PostResponse{
		ID:        post.ID,
		Content:   post.Content,
		Username:  post.User.Username,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	return resPost, nil
}

func (pu *postUsecase) GetPostByUserID(userId uint) (models.PostResponse, error) {
	posts := []models.Post{}
	if err := pu.pr.GetPostByUserID(&posts, userId); err != nil {
		return nil, err
	}
	resPosts := []models.PostResponse{}
	for _, v := range posts {
		p := models.PostResponse{
			ID:        v.ID,
			Content:   v.Content,
			Username:  v.User.Username,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resPosts = append(resPosts, p)
	}
	return resPosts, nil
}

func (pu *postUsecase) CreatePost(post models.Post) (models.PostResponse, error) {
	if err := pu.pv.PostValidate(post); err != nil {
		return models.PostResponse{}, err
	}
	if err := pu.pr.CreatePost(&post); err != nil {
		return models.PostResponse{}, err
	}
	resPost := models.PostResponse{
		ID:        post.ID,
		Content:   post.Content,
		Username:  post.User.Username,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	return resPost, nil
}

func (pu *postUsecase) UpdatePost(post models.Post, postId uint) (models.PostResponse, error) {
	if err := pu.pv.PostValidate(post); err != nil {
		return models.PostResponse{}, err
	}
	if err := pu.pr.UpdatePost(&post, postId); err != nil {
		return models.PostResponse{}, err
	}
	resPost := models.PostResponse{
		ID:        post.ID,
		Content:   post.Content,
		Username:  post.User.Username,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	return resPost, nil
}

func (pu *postUsecase) DeletePost(postId uint) error {
	if err := pu.pr.DeletePost(postId); err != nil {
		return err
	}
	return nil
}
