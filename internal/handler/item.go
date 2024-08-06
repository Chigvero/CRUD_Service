package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	todo "todo-app"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErroreResponse(c, 400, "invalid id param")
		return
	}
	var input todo.TodoItem
	err = c.BindJSON(&input)
	if err != nil {
		newErroreResponse(c, 400, err.Error())
		return
	}
	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error()) //500Error
		return
	}
	c.JSON(200, map[string]interface{}{
		"id": id,
	})
}

type getItemsResponse struct {
	Data []todo.TodoItem `json:"data"`
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErroreResponse(c, 400, "invalid list  id param")
		return
	}

	items, err := h.services.TodoItem.GetAllItems(userId, listId)
	if err != nil {

		newErroreResponse(c, 500, err.Error())
		return
	}
	c.JSON(200, getItemsResponse{
		Data: items,
	})
}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErroreResponse(c, 400, "invalid item id param")
		return
	}
	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErroreResponse(c, 500, err.Error())
		return
	}
	c.JSON(200, item)
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErroreResponse(c, 400, "invalid id param")
		return
	}
	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErroreResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = h.services.TodoItem.UpdateItem(userId, id, input); err != nil {
		newErroreResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(200, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErroreResponse(c, 400, "invalid item id param")
		return
	}
	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErroreResponse(c, 500, err.Error())
		return
	}
	c.JSON(200, map[string]interface{}{
		"StatusCode": "Ok",
	})
}
