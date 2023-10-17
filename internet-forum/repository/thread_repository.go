package repository

import (
	"github.com/junshintakeda/internet-forum/models"

	"gorm.io/gorm"
)

type IThreadRepository interface {
	GetAllThreads(threads *[]models.Thread) error
	GetThreadByID(thread *models.Thread, threadId uint) error
	CreateThread(thread *models.Thread) error
}

type threadrepository struct {
	db *gorm.DB
}

func NewThreadRepository(db *gorm.DB) IThreadRepository {
	return &threadrepository{db}
}

func (tr *threadrepository) GetAllThreads(threads *[]models.Thread) error {
	if err := tr.db.Order("created_at").Find(threads).Error; err != nil {
		return err
	}
	return nil
}

func (tr *threadrepository) GetThreadByID(thread *models.Thread, threadId uint) error {
	if err := tr.db.First(thread, threadId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *threadrepository) CreateThread(thread *models.Thread) error {
	if err := tr.db.Create(thread).Error; err != nil {
		return err
	}
	return nil
}
