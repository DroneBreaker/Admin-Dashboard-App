package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DroneBreaker/Admin-Dashboard-App/backend/internal/models"
	"github.com/DroneBreaker/Admin-Dashboard-App/backend/internal/service"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetAll(c echo.Context) error {
	users, err := h.userService.GetAll()
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) Create(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.userService.Create(*user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
		"user":    user,
	})
}

func (h *userHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := h.userService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "there is no such user",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("User with id %d found successfully", id),
		"user":    user,
	})
}

func (h *userHandler) GetByUsername(c echo.Context) error {
	username := c.Param("username")

	user, err := h.userService.GetByUsername(username)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "there is no such user",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Username found successfully",
		"user":    user,
	})
}
