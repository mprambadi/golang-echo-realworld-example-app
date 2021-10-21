package todo

import (
	"github.com/mprambadi/golang-echo-realworld-example-app/model"
)

type Store interface {
	GetByID(uint) (*model.Todo, error)
	Create(*model.Todo) error
	Update(*model.Todo) error
	List() ([]model.Todo, error)
	Delete(*model.Todo) error
}
