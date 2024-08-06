package service

import (
	todo "todo-app"
	"todo-app/internal/repository"
)

type TodoListService struct {
	repo repository.ToDoList
}

func NewTodoListService(repos repository.ToDoList) *TodoListService {
	return &TodoListService{
		repo: repos,
	}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
func (s *TodoListService) GetAllLists(userId int) ([]todo.TodoList, error) {

	return s.repo.GetAllLists(userId)
}

func (s *TodoListService) GetListById(userId, id int) (todo.TodoList, error) {
	return s.repo.GetListById(userId, id)
}

func (s *TodoListService) DeleteById(userId, id int) error {
	return s.repo.DeleteById(userId, id)
}

func (s *TodoListService) UpdateList(userId, id int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateList(userId, id, input)
}
