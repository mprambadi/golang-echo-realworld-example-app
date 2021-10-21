package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mprambadi/golang-echo-realworld-example-app/model"
	"github.com/mprambadi/golang-echo-realworld-example-app/utils"
)

// GetById godoc
// @Summary Get Todo By Id
// @Description Get todo by id
// @ID get-todo
// @Accept  json
// @Produce  json
// @Tags todo
// @Param id path string true "id of the todo to get"
// @Success 200 {object} todoResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /todos/{id} [get]
func (h *Handler) GetTodoById(c echo.Context) error {
	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		todoId = 0
	}

	a, err := h.todoStore.GetByID(uint(todoId))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	return c.JSON(http.StatusOK, newTodoResponse(c, a))
}

// GetTodoList godoc
// @Summary Get Todo List
// @Description Get todo List
// @Accept  json
// @Produce  json
// @ID todos
// @Tags todo
// @Success 200 {object} todoListResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /todos [get]
func (h *Handler) GetTodoList(c echo.Context) error {

	todos, err := h.todoStore.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newTodoListResponse(todos))
}

// CreateTodo godoc
// @Summary Create an todo
// @Description Create an todo. Auth is require
// @ID create-todo
// @Tags todo
// @Accept  json
// @Produce  json
// @Param todo body todoCreateRequest true "Todo to create"
// @Success 201 {object} singleTodoResponse
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /todos [post]
func (h *Handler) CreateTodo(c echo.Context) error {
	var a model.Todo

	req := &todoCreateRequest{}
	if err := req.bind(c, &a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	a.AuthorID = userIDFromToken(c)

	err := h.todoStore.Create(&a)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusCreated, newTodoResponse(c, &a))
}

// Deletetodo godoc
// @Summary Delete an todo
// @Description Delete an todo. Auth is required
// @ID delete-todo
// @Tags todo
// @Accept  json
// @Produce  json
// @Param id path string true "id of the todo to delete"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /todos/{id} [delete]
func (h *Handler) DeleteTodo(c echo.Context) error {
	todoId, err := strconv.Atoi(c.Param("id"))

	a, err := h.todoStore.GetByID(uint(todoId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	err = h.todoStore.Delete(a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}

// UpdateTodo godoc
// @Summary Update an todo
// @Description Update an todo. Auth is required
// @ID update-todo
// @Tags todo
// @Accept  json
// @Produce  json
// @Param id path string true "id of the todo to update"
// @Param todo body todoUpdateRequest true "Todo to update"
// @Success 200 {object} singleTodoResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /todos/{id} [put]
func (h *Handler) UpdateTodo(c echo.Context) error {
	todoId, err := strconv.Atoi(c.Param("id"))

	a, err := h.todoStore.GetByID(uint(todoId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	req := &todoUpdateRequest{}

	if err := req.bind(c, a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err = h.todoStore.Update(a); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, newTodoResponse(c, a))
}
