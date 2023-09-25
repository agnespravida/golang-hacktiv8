package routes

import (
	"final-project/controllers"
	"final-project/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.Register)
		userRouter.POST("/login", controllers.Login)
	}

	socialMediaRouter := router.Group("/social-media")

	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.SocialMediaCreate)
		socialMediaRouter.GET("/all", controllers.SocialMediaGetAll)
		socialMediaRouter.GET("/user-login", controllers.SocialMediaGetByUserLogin)
		socialMediaRouter.GET("/:socmedID", controllers.SocialMediaGetByID)
		socialMediaRouter.PUT("/:socmedID", middlewares.SocialMediaAuth(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socmedID", middlewares.SocialMediaAuth(), controllers.DeleteSocialMedia)
	}

	PhotoRouter := router.Group("/photo")

	{
		PhotoRouter.Use(middlewares.Authentication())
		PhotoRouter.POST("/", controllers.PhotoCreate)
		PhotoRouter.GET("/all", controllers.PhotoGetAll)
		PhotoRouter.GET("/user-login", controllers.PhotoGetByUserLogin)
		PhotoRouter.GET("/:photoID", controllers.PhotoGetByID)
		PhotoRouter.PUT("/:photoID", middlewares.PhotoAuth(), controllers.UpdatePhoto)
		PhotoRouter.DELETE("/:photoID", middlewares.PhotoAuth(), controllers.DeletePhoto)
	}

	commentRouter := router.Group("/photo")

	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.CommentCreate)
		commentRouter.GET("/all", controllers.CommentGetAll)
		commentRouter.GET("/user-login", controllers.CommentGetByUserLogin)
		commentRouter.GET("/:commentID", controllers.CommentGetByID)
		commentRouter.PUT("/:commentID", middlewares.CommentAuth(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentID", middlewares.CommentAuth(), controllers.DeleteComment)
	}

	// router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
