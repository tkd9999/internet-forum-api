package main

import (
	"github.com/junshintakeda/internet-forum/controller"
	"github.com/junshintakeda/internet-forum/db"
	"github.com/junshintakeda/internet-forum/repository"
	"github.com/junshintakeda/internet-forum/router"
	"github.com/junshintakeda/internet-forum/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	threadRepository := repository.NewThreadRepository(db)
	postRepository := repository.NewPostRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	threadUsecase := usecase.NewThreadUsecase(threadRepository)
	postUsecase := usecase.NewPostUsecase(postRepository)
	userController := controller.NewUserController(userUsecase)
	threadController := controller.NewThreadController(threadUsecase)
	postController := controller.NewPostController(postUsecase, threadUsecase)
	e := router.NewRouter(userController, threadController, postController)
	e.Logger.Fatal(e.Start(":8080"))
}
