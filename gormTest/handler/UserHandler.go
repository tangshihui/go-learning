package handler

import (
	"example/go-learning/gormTest/model"
	"example/go-learning/gormTest/repo"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandler struct {
	repo *repo.UserRepo
}

func NewUserHandler(repo *repo.UserRepo) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) ListAllUsers(e echo.Context) error {
	page, _ := strconv.Atoi(e.QueryParam("page"))
	perPage, _ := strconv.Atoi(e.QueryParam("perPage"))
	users, _ := h.repo.ListAllUsers(page, perPage)
	return e.JSON(http.StatusOK, users)
}

// SaveUser 新增用户
// @Tags user
// @Summary 新增用户
// @Description 新增用户
// @ID api-user-update
// @Accept  json
// @Produce  json
// @Param body body model.User true "user 数据"
// @Success 200 {object} model.User
// @Router /user [put]
func (h *UserHandler) SaveUser(e echo.Context) error {
	user := model.User{}
	e.Bind(&user)

	saveUser, _ := h.repo.SaveUser(user)
	return e.JSON(http.StatusOK, saveUser)
}

func (h *UserHandler) DeleteUser(e echo.Context) error {
	id := e.Param("id")
	idV, _ := strconv.Atoi(id)
	users, _ := h.repo.DeleteUser(idV)
	return e.JSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(e echo.Context) error {
	user := model.User{}
	e.Bind(&user)
	rows, _ := h.repo.UpdateUser(&user)
	return e.JSON(http.StatusOK, rows)
}

func (h *UserHandler) QueryUser(e echo.Context) error {
	id := e.Param("id")
	idV, _ := strconv.Atoi(id)
	user, err := h.repo.QueryUser(idV)
	if err != nil {
		res := make(map[string]string, 0)
		res["message"] = err.Error()
		return e.JSON(http.StatusOK, res)
	}
	return e.JSON(http.StatusOK, user)
}
