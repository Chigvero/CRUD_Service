package repository

import (
	"github.com/jmoiron/sqlx"
	todo "todo-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type ToDoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAllLists(userId int) ([]todo.TodoList, error)
	GetListById(userId, id int) (todo.TodoList, error)
	DeleteById(userId, id int) error
	UpdateList(userId, id int, input todo.UpdateListInput) error
}

type ToDoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAllItems(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	UpdateItem(userId, itemId int, input todo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ToDoList:      NewTodoListPostgres(db),
		ToDoItem:      NewTodoItemPostgres(db),
	}
}
