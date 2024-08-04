package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "todo-app"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{
		db: db,
	}
}

func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s(user_id,list_id)VALUES($1,$2) ", usersListsTable)
	_, err = tx.Exec(createUsersListsQuery, userId, id)
	if err != nil {
		tx.Rollback()
		fmt.Println(2)
		return 0, err
	}
	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAllListsM1(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	var listId int
	getUserListsquery := fmt.Sprintf("SELECT list_id FROM %s WHERE user_id=$1", usersListsTable)
	rows, err := tx.Query(getUserListsquery, userId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for rows.Next() {
		var list1 todo.TodoList
		rows.Scan(&listId)
		getTodoListQuery := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", todoListsTable)
		err = tx.QueryRow(getTodoListQuery, listId).Scan(&list1.Id, &list1.Title, &list1.Description)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		lists = append(lists, list1)
	}
	return lists, nil
}
func (r *TodoListPostgres) GetAllLists(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := fmt.Sprintf("SELECT tl.id , tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id=ul.list_id where ul.user_id=$1",
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *TodoListPostgres) GetListById(userId, id int) (todo.TodoList, error) {
	var list todo.TodoList
	query := fmt.Sprintf("SELECT tl.id , tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id=ul.list_id where ul.user_id=$1 and tl.id=$2",
		todoListsTable, usersListsTable)

	err := r.db.Get(&list, query, userId, id)
	return list, err
}