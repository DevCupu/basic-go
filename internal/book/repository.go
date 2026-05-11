package book

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Create(book *Book) error
	FindAll() ([]Book, error)
	FindByID(id uint) (*Book, error)
	Update(book *Book) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(book *Book) error {
	return r.db.Create(book).Error
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) FindByID(id uint) (*Book, error) {
	var book Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, fmt.Errorf("buku dengan ID %d tidak ditemukan", id)
	}
	return &book, nil
}

func (r *repository) Update(book *Book) error {
	return r.db.Save(book).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Book{}, id).Error
}
