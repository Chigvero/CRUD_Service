package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	todo "todo-app"
)

func (h *Handler) signUP(c *gin.Context) {
	user := todo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		newErroreResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateUser(user)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input SignInInput
	err := c.BindJSON(&input)
	if err != nil {
		newErroreResponse(c, 400, err.Error())
		return
	}
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErroreResponse(c, 401, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
