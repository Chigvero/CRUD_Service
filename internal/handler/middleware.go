package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

// Нам нужно получать значение из Хедера авторизации
//
//	Валидировать его, парсить токен и записывать пользователя в контекст
func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErroreResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErroreResponse(c, 401, "invalid auth header")
		return
	}

	//Parse token
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErroreResponse(c, 401, err.Error())
		return
	}
	c.Set("userId", userId)
}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErroreResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	IdInt, ok := id.(int)
	if !ok {
		newErroreResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}
	return IdInt, nil
}
