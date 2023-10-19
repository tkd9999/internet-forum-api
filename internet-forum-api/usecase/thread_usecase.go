package usecase

import (
	"github.com/junshintakeda/internet-forum/models"
	"github.com/junshintakeda/internet-forum/repository"
	"github.com/junshintakeda/internet-forum/validator"
)

type IThreadUsecase interface {
	GetAllThreads() ([]models.ThreadResponse, error)
	CreateThread(thread models.Thread) (models.ThreadResponse, error)
}

type threadUsecase struct {
	tr repository.IThreadRepository
	tv validator.IThreadValidator
}

func NewThreadUsecase(tr repository.IThreadRepository, tv validator.IThreadValidator) IThreadUsecase {
	return &threadUsecase{tr, tv}
}

func (tu *threadUsecase) GetAllThreads() ([]models.ThreadResponse, error) {
	threads := []models.Thread{}
	if err := tu.tr.GetAllThreads(&threads); err != nil {
		return nil, err
	}
	resThreads := []models.ThreadResponse{}
	for _, v := range threads {
		t := models.ThreadResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resThreads = append(resThreads, t)
	}
	return resThreads, nil
}

func (tu *threadUsecase) CreateThread(thread models.Thread) (models.ThreadResponse, error) {
	if err := tu.tv.ThreadValidate(thread); err != nil {
		return models.ThreadResponse{}, err
	}
	if err := tu.tr.CreateThread(&thread); err != nil {
		return models.ThreadResponse{}, err
	}
	resThread := models.ThreadResponse{
		ID:        thread.ID,
		Title:     thread.Title,
		CreatedAt: thread.CreatedAt,
		UpdatedAt: thread.UpdatedAt,
	}
	return resThread, nil
}
