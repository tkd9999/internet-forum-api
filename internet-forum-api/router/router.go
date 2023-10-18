package router

import (
	"net/http"
	"os"
	"practice/go-rest-api/controller"

	"github.com/junshintakeda/internet-forum-api/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.IThreadController, pc controller.IPostController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
		// CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge: 60,
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	t := e.Group("/threads")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllThreads)
	t.GET("/:threadId", tc.GetThreadByID)
	t.GET("/user/:userId", tc.GetThreadsByUserID)
	t.POST("", tc.CreateThread)
	t.PUT("/:threadId", tc.UpdateThread)
	t.DELETE("/:threadId", tc.DeleteThread)

	p := e.Group("/posts")
	p.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	p.GET("/:threadId", pc.GetAllPosts)
	p.GET("/post/:postId", pc.GetPostByID)
	p.GET("/user/:userId", pc.GetPostsByUserID)
	p.POST("/:threadId", pc.CreatePost)
	p.PUT("/:postId", pc.UpdatePost)
	p.DELETE("/:postId", pc.DeletePost)
	return e
}
