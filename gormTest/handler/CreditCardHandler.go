package handler

import (
	"example/go-learning/gormTest/model"
	"example/go-learning/gormTest/repo"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CreditCardHandler struct {
	repo *repo.CreditCardRepo
}

func NewCreditCardHandler(repo *repo.CreditCardRepo) *CreditCardHandler {
	return &CreditCardHandler{repo: repo}
}

func (h *CreditCardHandler) ListAll(e echo.Context) error {
	page, _ := strconv.Atoi(e.QueryParam("page"))
	perPage, _ := strconv.Atoi(e.QueryParam("perPage"))
	users, _ := h.repo.ListAll(page, perPage)
	return e.JSON(http.StatusOK, users)
}

func (h *CreditCardHandler) SaveCreditCard(e echo.Context) error {
	card := model.CreditCard{}
	e.Bind(&card)

	savedCard, _ := h.repo.SaveCreditCard(card)
	return e.JSON(http.StatusOK, savedCard)
}

//func (h *UserHandler) DeleteUser(e echo.Context) error {
//	id := e.Param("id")
//	idV, _ := strconv.Atoi(id)
//	users, _ := h.repo.DeleteUser(idV)
//	return e.JSON(http.StatusOK, users)
//}
//
//func (h *UserHandler) UpdateUser(e echo.Context) error {
//	user := model.User{}
//	e.Bind(&user)
//	rows, _ := h.repo.UpdateUser(&user)
//	return e.JSON(http.StatusOK, rows)
//}

func (h *CreditCardHandler) QueryById(e echo.Context) error {
	id := e.Param("id")
	idV, _ := strconv.Atoi(id)
	card, err := h.repo.QueryById(idV)
	if err != nil {
		res := make(map[string]string, 0)
		res["message"] = err.Error()
		return e.JSON(http.StatusOK, res)
	}
	return e.JSON(http.StatusOK, card)
}

func (h *CreditCardHandler) QueryByUserId(e echo.Context) error {
	id := e.Param("userId")
	idV, _ := strconv.Atoi(id)
	cards, err := h.repo.QueryByUserId(idV)
	if err != nil {
		res := make(map[string]string, 0)
		res["message"] = err.Error()
		return e.JSON(http.StatusOK, res)
	}
	return e.JSON(http.StatusOK, cards)
}
