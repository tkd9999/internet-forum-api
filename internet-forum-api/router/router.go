package router

import (
	"os"

	"github.com/junshintakeda/internet-forum/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, tc controller.IThreadController, pc controller.IPostController) *echo.Echo {
	e := echo.New()

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)

	t := e.Group("/threads")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllThreads)
	t.POST("", tc.CreateThread)
	t.GET("/:threadId", pc.GetAllPosts)
	t.POST("/:threadId", pc.CreatePost)

	p := e.Group("/posts")
	p.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	p.GET("/:postId", pc.GetPostByID)
	p.GET("/user/:userId", pc.GetPostsByUserID)
	p.PUT("/:postId", pc.UpdatePost)
	p.DELETE("/:postId", pc.DeletePost)
	return e
}
