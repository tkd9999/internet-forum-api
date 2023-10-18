package controller

import (
	"net/http"

	"github.com/junshintakeda/internet-forum/models"
	"github.com/junshintakeda/internet-forum/usecase"

	"github.com/labstack/echo/v4"
)

type IPostController interface {
	GetAllPosts(c echo.Context) error
	GetPostByID(c echo.Context) error
	GetPostsByUserID(c echo.Context) error
	CreatePost(c echo.Context) error
	UpdatePost(c echo.Context) error
	DeletePost(c echo.Context) error
}

type postController struct {
	pu usecase.IPostUsecase
}

func NewPostController(pu usecase.IPostUsecase) IPostController {
	return &postController{pu}
}

func (pc *postController) GetAllPosts(c echo.Context) error {
	threadId := c.Param("threadId")
	threads, err := pc.pu.GetAllPosts(threadId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, threads)
}

func (pc *postController) GetPostByID(c echo.Context) error {
	postId := c.Param("postId")
	post, err := pc.pu.GetPostByID(postId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, post)
}

func (pc *postController) GetPostsByUserID(c echo.Context) error {
	userId := c.Param("userId")
	posts, err := pc.pu.GetPostsByUserID(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, posts)
}

func (pc *postController) CreatePost(c echo.Context) error {
	threadId := c.Param("threadId")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	post := models.Post{}
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	post.UserId = uint(userId.(float64))
	post.ThreadId = uint(threadId)
	postRes, err := pc.pu.CreatePost(post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, postRes)
}

func (pc *postController) UpdatePost(c echo.Context) error {
	postId := c.Param("postId")
	post := models.Post{}
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	postRes, err := pc.pu.UpdatePost(postId, post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postRes)
}

func (pc *postController) DeletePost(c echo.Context) error {
	postId := c.Param("postId")
	if err := pc.pu.DeletePost(postId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
