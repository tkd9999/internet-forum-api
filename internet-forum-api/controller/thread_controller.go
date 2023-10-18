package controller

import (
	"net/http"

	"github.com/junshintakeda/internet-forum/models"
	"github.com/junshintakeda/internet-forum/usecase"

	"github.com/labstack/echo/v4"
)

type IThreadController interface {
	GetAllThreads(c echo.Context) error
	GetThreadByID(c echo.Context) error
	CreateThread(c echo.Context) error
}

type threadController struct {
	tc usecase.IThreadUsecase
}

func NewThreadController(tc usecase.IThreadUsecase) IThreadController {
	return &threadController{tc}
}

func (tc *threadController) GetAllThreads(c echo.Context) error {
	threads, err := tc.tc.GetAllThreads()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, threads)
}

func (tc *threadController) GetThreadByID(c echo.Context) error {
	threadId := c.Param("threadId")
	thread, err := tc.tc.GetThreadByID(threadId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, thread)
}

func (tc *threadController) CreateThread(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	thread := models.Thread{}
	if err := c.Bind(&thread); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	thread.UserId = uint(userId.(float64))
	threadRes, err := tc.tc.CreateThread(thread)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, threadRes)
}
