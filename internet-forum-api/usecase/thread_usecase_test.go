package usecase

import (
	"fmt"
	"testing"
	"time"

	"github.com/junshintakeda/internet-forum/models"
)

type MockThreadRepository struct {
	GetAllThreadsResponse []models.Thread
	GetThreadByIDResponse models.Thread
	CreateThreadError     error
}

func (m *MockThreadRepository) GetAllThreads(threads *[]models.Thread) error {
	*threads = m.GetAllThreadsResponse
	return nil
}

func (m *MockThreadRepository) GetThreadByID(thread *models.Thread, threadID uint) error {
	*thread = m.GetThreadByIDResponse
	return nil
}

func (m *MockThreadRepository) CreateThread(thread *models.Thread) error {
	return m.CreateThreadError
}

type MockThreadValidator struct {
	ThreadValidateError error
}

func (m *MockThreadValidator) ThreadValidate(thread models.Thread) error {
	return m.ThreadValidateError
}

func TestThreadUsecase_GetAllThreads(t *testing.T) {
	mockthreadRepository := &MockThreadRepository{}
	mockthreadValidator := &MockThreadValidator{}

	tu := NewThreadUsecase(mockthreadRepository, mockthreadValidator)

	threads := []models.Thread{
		{
			ID:        1,
			Title:     "Sample Thread 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     "Sample Thread 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	mockthreadRepository.GetAllThreadsResponse = threads

	resThreads, err := tu.GetAllThreads()

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if len(resThreads) != len(threads) {
		t.Errorf("Expected %d threads, but got %d", len(threads), len(resThreads))
	}

	for i := range threads {
		if threads[i].ID != resThreads[i].ID ||
			threads[i].Title != resThreads[i].Title ||
			threads[i].CreatedAt != resThreads[i].CreatedAt ||
			threads[i].UpdatedAt != resThreads[i].UpdatedAt {
			t.Errorf("Thread %d details do not match", threads[i].ID)
		}
	}
}

func TestThreadUsecase_GetThreadByID(t *testing.T) {
	mockthreadRepository := &MockThreadRepository{}
	mockthreadValidator := &MockThreadValidator{}

	tu := NewThreadUsecase(mockthreadRepository, mockthreadValidator)

	thread := models.Thread{
		ID:        1,
		Title:     "Sample Thread 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockthreadRepository.GetThreadByIDResponse = thread

	threadId := thread.ID
	resThread, err := tu.GetThreadByID(threadId)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if thread.ID != resThread.ID ||
		thread.Title != resThread.Title ||
		thread.CreatedAt != resThread.CreatedAt ||
		thread.UpdatedAt != resThread.UpdatedAt {
		t.Error("Thread details do not match")
	}

}

func TestThreadUsecase_CreateThread(t *testing.T) {
	mockthreadRepository := &MockThreadRepository{}
	mockthreadValidator := &MockThreadValidator{}

	tu := NewThreadUsecase(mockthreadRepository, mockthreadValidator)

	thread := models.Thread{
		Title: "Sample Thread 1",
	}

	mockthreadRepository.CreateThreadError = nil
	mockthreadValidator.ThreadValidateError = nil

	resThread, err := tu.CreateThread(thread)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if resThread.Title != thread.Title {
		t.Errorf("Expected %s, but got %s", thread.Title, resThread.Title)
	}
}

func TestThreadUsecase_CreateThread_ValidationError(t *testing.T) {
	mockthreadRepository := &MockThreadRepository{}
	mockthreadValidator := &MockThreadValidator{}

	tu := NewThreadUsecase(mockthreadRepository, mockthreadValidator)

	thread := models.Thread{
		Title: "",
	}

	mockthreadRepository.CreateThreadError = nil
	mockthreadValidator.ThreadValidateError = fmt.Errorf("Title is required")

	_, err := tu.CreateThread(thread)

	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
