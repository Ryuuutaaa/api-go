package controllers

import (
	"demo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
)

func Create(c echo.Context) error {
	ctx := c.Request().Context()

	policy := bluemonday.UGCPolicy()

	name := c.FormValue("name")
	email := c.FormValue("email")

	cleanName := policy.Sanitize(name)
	cleanEmail := policy.Sanitize(email)

	if name == "" || email == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Bad Request",
		})
	}

	data := models.User{
		Name:  cleanName,
		Email: cleanEmail,
	}

	if err := models.Create(ctx, data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "Success",
		"data":    data,
	})
}

func Read(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Bad Request",
		})
	}

	user, err := models.Read(ctx, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Success",
		"data":    user,
	})
}

func Update(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Bad Request",
		})
	}

	name := c.FormValue("name")
	email := c.FormValue("email")

	cleanName := bluemonday.UGCPolicy().Sanitize(name)
	cleanEmail := bluemonday.UGCPolicy().Sanitize(email)

	if name == "" || email == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Bad Request",
		})
	}

	data := models.User{
		Name:  cleanName,
		Email: cleanEmail,
	}

	user, err := models.Update(ctx, uint(id), data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Success",
		"data":    user,
	})
}

func Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "ID is required",
		})
	}

	if err := models.Delete(ctx, id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]any{
				"message": "User not found",
				"id":      id,
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "User deleted successfully",
	})
}
