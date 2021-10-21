package store

import (
	"github.com/jinzhu/gorm"
	"github.com/mprambadi/golang-echo-realworld-example-app/model"
)

type TodoStore struct {
	db *gorm.DB
}

func NewTodoStore(db *gorm.DB) *TodoStore {
	return &TodoStore{
		db: db,
	}
}

func (us *TodoStore) GetByID(id uint) (*model.Todo, error) {
	var m model.Todo
	if err := us.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *TodoStore) Create(u *model.Todo) (err error) {
	return us.db.Create(u).Error
}

func (us *TodoStore) Update(u *model.Todo) error {
	return us.db.Model(u).Update(u).Error
}

func (as *TodoStore) List() ([]model.Todo, error) {
	var todos []model.Todo
	if err := as.db.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (as *TodoStore) Delete(a *model.Todo) error {
	return as.db.Delete(a).Error
}
