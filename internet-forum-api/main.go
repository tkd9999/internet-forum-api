package main

import (
	"github.com/junshintakeda/internet-forum/controller"
	"github.com/junshintakeda/internet-forum/db"
	"github.com/junshintakeda/internet-forum/repository"
	"github.com/junshintakeda/internet-forum/router"
	"github.com/junshintakeda/internet-forum/usecase"
	"github.com/junshintakeda/internet-forum/validator"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	threadRepository := repository.NewThreadRepository(db)
	postRepository := repository.NewPostRepository(db)
	userValidator := validator.NewUserValidator()
	threadValidator := validator.NewThreadValidator()
	postValidator := validator.NewPostValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	threadUsecase := usecase.NewThreadUsecase(threadRepository, threadValidator)
	postUsecase := usecase.NewPostUsecase(postRepository, postValidator)
	userController := controller.NewUserController(userUsecase)
	threadController := controller.NewThreadController(threadUsecase)
	postController := controller.NewPostController(postUsecase, threadUsecase)
	e := router.NewRouter(userController, threadController, postController)
	e.Logger.Fatal(e.Start(":8080"))
}
