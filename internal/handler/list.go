package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	todo "todo-app"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}
	var input todo.TodoList
	err = c.BindJSON(&input)
	if err != nil {
		newErroreResponse(c, 400, err.Error())
		return
	}

	//call service method
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}
	lists, err := h.services.TodoList.GetAllLists(userId)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErroreResponse(c, 400, "invalid id param")
		return
	}
	list, err := h.services.TodoList.GetListById(userId, id)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, list)
}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
