package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	todo "todo-app"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{
		db: db,
	}
}

func (r *TodoItemPostgres) Create(listId int, item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title,description) VALUES($1,$2) RETURNING id", todoItemsTable)
	err = tx.QueryRow(createItemQuery, item.Title, item.Description).Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id,item_id) VALUES($1,$2)", ListsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return itemId, tx.Commit()
}

func (r *TodoItemPostgres) GetAllItems(userId, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti inner JOIN %s li on ti.id=li.item_id "+
		"INNER JOIN %s ul on ul.list_id=li.list_id WHERE li.list_id=$1 AND ul.user_id=$2",
		todoItemsTable, ListsItemsTable, usersListsTable)
	rows, err := r.db.Query(query, listId, userId)
	if err != nil {
		fmt.Println(1)
		return nil, err
	}
	for rows.Next() {
		var item todo.TodoItem
		err = rows.Scan(&item.Id, &item.Title, &item.Description, &item.Done)
		if err != nil {
			fmt.Println(2)
			return nil, err
		}
		items = append(items, item)
	}
	//if err = r.db.Select(&items, query, listId, userId); err != nil {
	//	return nil, err
	//}надо разобраться)
	return items, nil
}

func (r *TodoItemPostgres) GetById(userId, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti inner JOIN %s li on ti.id=li.item_id "+
		"INNER JOIN %s ul on ul.list_id=li.list_id WHERE ti.id=$1 AND ul.user_id=$2",
		todoItemsTable, ListsItemsTable, usersListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s li , %s ul WHERE ti.id=li.item_id AND li.list_id=ul.list_id AND ul.user_id=$1 AND ti.id=$2",
		todoItemsTable, ListsItemsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *TodoItemPostgres) UpdateItem(userId, itemId int, input todo.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}
	//title=$1
	//description=$1
	//title=$1, description=$2

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id=li.item_id AND li.list_id=ul.list_id AND ul.user_id=$%d AND ti.id=$%d",
		todoItemsTable, setQuery, ListsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, userId, itemId)
	_, err := r.db.Exec(query, args...)
	return err
}
