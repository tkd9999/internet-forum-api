package controller

import (
	"net/http"
	"strconv"

	"github.com/junshintakeda/internet-forum/models"
	"github.com/junshintakeda/internet-forum/usecase"

	"github.com/golang-jwt/jwt/v4"
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
	tu usecase.IThreadUsecase
}

func NewPostController(pu usecase.IPostUsecase, tu usecase.IThreadUsecase) IPostController {
	return &postController{pu, tu}
}

func (pc *postController) GetAllPosts(c echo.Context) error {
	ThreadId := c.Param("threadId")
	threadId, _ := strconv.Atoi(ThreadId)
	posts, err := pc.pu.GetAllPosts(uint(threadId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	thread, err := pc.tu.GetThreadByID(uint(threadId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"posts": posts, "thread": thread})
}

func (pc *postController) GetPostByID(c echo.Context) error {
	PostId := c.Param("postId")
	postId, _ := strconv.Atoi(PostId)
	post, err := pc.pu.GetPostByID(uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, post)
}

func (pc *postController) GetPostsByUserID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	posts, err := pc.pu.GetPostByUserID(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, posts)
}

func (pc *postController) CreatePost(c echo.Context) error {
	ThreadId := c.Param("threadId")
	threadId, _ := strconv.Atoi(ThreadId)
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
	PostId := c.Param("postId")
	postId, _ := strconv.Atoi(PostId)
	post := models.Post{}
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	postRes, err := pc.pu.UpdatePost(post, uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postRes)
}

func (pc *postController) DeletePost(c echo.Context) error {
	PostId := c.Param("postId")
	postId, _ := strconv.Atoi(PostId)
	if err := pc.pu.DeletePost(uint(postId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
