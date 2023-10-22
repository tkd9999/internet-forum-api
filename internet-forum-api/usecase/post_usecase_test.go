package usecase

import (
	"github.com/junshintakeda/internet-forum/models"

	"testing"
	"time"
)

type MockPostRepository struct {
	GetAllPostsResponse     []models.Post
	GetPostByIDResponse     models.Post
	GetPostByUserIDResponse []models.Post
	CreatePostError         error
	UpdatePostError         error
	DeletePostError         error
}

func (m *MockPostRepository) GetAllPosts(posts *[]models.Post, threadId uint) error {
	*posts = m.GetAllPostsResponse
	return nil
}

func (m *MockPostRepository) GetPostByID(post *models.Post, postId uint) error {
	*post = m.GetPostByIDResponse
	return nil
}

func (m *MockPostRepository) GetPostByUserID(posts *[]models.Post, userId uint) error {
	*posts = m.GetPostByUserIDResponse
	return nil
}

func (m *MockPostRepository) CreatePost(post *models.Post) error {
	return m.CreatePostError
}

func (m *MockPostRepository) UpdatePost(post *models.Post, postId uint) error {
	return m.UpdatePostError
}

func (m *MockPostRepository) DeletePost(postId uint) error {
	return m.DeletePostError
}

type MockPostValidator struct {
	PostValidateError error
}

func (m *MockPostValidator) PostValidate(post models.Post) error {
	return m.PostValidateError
}

func TestGetAllPosts(t *testing.T) {
	MockPostRepository := &MockPostRepository{}
	MockPostValidator := &MockPostValidator{}

	pu := NewPostUsecase(MockPostRepository, MockPostValidator)
	posts := []models.Post{
		{
			ID:        1,
			Content:   "Sample Post 1",
			UserId:    1,
			ThreadId:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Content:   "Sample Post 2",
			UserId:    2,
			ThreadId:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	MockPostRepository.GetAllPostsResponse = posts

	resPosts, err := pu.GetAllPosts(1)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if len(resPosts) != len(posts) {
		t.Errorf("Expected %d threads, but got %d", len(posts), len(resPosts))
	}

	for i, v := range resPosts {
		if posts[i].ID != v.ID ||
			posts[i].Content != v.Content ||
			posts[i].CreatedAt != v.CreatedAt ||
			posts[i].UpdatedAt != v.UpdatedAt {
			t.Errorf("Thread %d details do not match", posts[i].ID)
		}
	}
}

func TestGetPostByID(t *testing.T) {
	MockPostRepository := &MockPostRepository{}
	MockPostValidator := &MockPostValidator{}

	pu := NewPostUsecase(MockPostRepository, MockPostValidator)

	post := models.Post{
		ID:        1,
		Content:   "Sample Post 1",
		UserId:    1,
		ThreadId:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	MockPostRepository.GetPostByIDResponse = post

	resPost, err := pu.GetPostByID(1)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if post.ID != resPost.ID ||
		post.Content != resPost.Content ||
		post.CreatedAt != resPost.CreatedAt ||
		post.UpdatedAt != resPost.UpdatedAt {
		t.Errorf("Post %d details do not match", post.ID)
	}
}

func TestGetPostByUserID(t *testing.T) {
	MockPostRepository := &MockPostRepository{}
	MockPostValidator := &MockPostValidator{}

	pu := NewPostUsecase(MockPostRepository, MockPostValidator)
	posts := []models.Post{
		{
			ID:        1,
			Content:   "Sample Post 1",
			UserId:    1,
			ThreadId:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Content:   "Sample Post 2",
			UserId:    1,
			ThreadId:  2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	MockPostRepository.GetPostByUserIDResponse = posts

	resPosts, err := pu.GetPostByUserID(1)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if len(resPosts) != len(posts) {
		t.Errorf("Expected %d threads, but got %d", len(posts), len(resPosts))
	}

	for i, v := range resPosts {
		if posts[i].ID != v.ID ||
			posts[i].Content != v.Content ||
			posts[i].CreatedAt != v.CreatedAt ||
			posts[i].UpdatedAt != v.UpdatedAt {
			t.Errorf("Thread %d details do not match", posts[i].ID)
		}
	}
}

func TestCreatePost(t *testing.T) {
	MockPostRepository := &MockPostRepository{}
	MockPostValidator := &MockPostValidator{}

	pu := NewPostUsecase(MockPostRepository, MockPostValidator)

	post := models.Post{
		ID:        1,
		Content:   "Sample Post 1",
		UserId:    1,
		ThreadId:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	resPost, err := pu.CreatePost(post)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if post.ID != resPost.ID ||
		post.Content != resPost.Content ||
		post.CreatedAt != resPost.CreatedAt ||
		post.UpdatedAt != resPost.UpdatedAt {
		t.Errorf("Post %d details do not match", post.ID)
	}
}

func TestUpdatePost(t *testing.T) {
	MockPostRepository := &MockPostRepository{}
	MockPostValidator := &MockPostValidator{}

	pu := NewPostUsecase(MockPostRepository, MockPostValidator)

	post := models.Post{
		ID:        1,
		Content:   "Sample Post 1",
		UserId:    1,
		ThreadId:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	resPost, err := pu.UpdatePost(post, 1)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if post.ID != resPost.ID ||
		post.Content != resPost.Content ||
		post.CreatedAt != resPost.CreatedAt ||
		post.UpdatedAt != resPost.UpdatedAt {
		t.Errorf("Post %d details do not match", post.ID)
	}
}

func TestDeletePost(t *testing.T) {
	MockPostRepository := &MockPostRepository{}
	MockPostValidator := &MockPostValidator{}

	pu := NewPostUsecase(MockPostRepository, MockPostValidator)

	err := pu.DeletePost(1)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
