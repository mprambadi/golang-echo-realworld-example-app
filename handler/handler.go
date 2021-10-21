package handler

import (
	"github.com/mprambadi/golang-echo-realworld-example-app/article"
	"github.com/mprambadi/golang-echo-realworld-example-app/todo"
	"github.com/mprambadi/golang-echo-realworld-example-app/user"
)

type Handler struct {
	userStore    user.Store
	articleStore article.Store
	todoStore    todo.Store
}

func NewHandler(us user.Store, as article.Store, ts todo.Store) *Handler {
	return &Handler{
		userStore:    us,
		articleStore: as,
		todoStore:    ts,
	}
}
