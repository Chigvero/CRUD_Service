package service

import (
	todo "todo-app"
	"todo-app/internal/repository"
)

type TodoItemService struct {
	repo     repository.ToDoItem
	listRepo repository.ToDoList
}

func NewTodoItemService(repo repository.ToDoItem, listRepo repository.ToDoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil {
		//list does not exist
		return 0, err
	}
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAllItems(userId, listId int) ([]todo.TodoItem, error) {
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil {
		//list does not exist
		return nil, err

	}
	return s.repo.GetAllItems(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}
func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}
func (s *TodoItemService) UpdateItem(userId, itemId int, input todo.UpdateItemInput) error {
	return s.repo.UpdateItem(userId, itemId, input)
}
