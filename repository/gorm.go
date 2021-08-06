package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"todo/entities"
)

type GormInsert struct {
	db *gorm.DB
}

func (insert GormInsert) Insert(v interface{}) error {
	return errors.WithMessage(insert.db.Create(v).Error,"gorm insert")
}

func (insert GormInsert) NewTask(task *entities.Task) error {
	return errors.WithMessage(insert.db.Create(v).Error,"gorm insert")
}