package router

import (
	"blog-system/handlers"
	"blog-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.ErrorHandler())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Auth routes (public)
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	// Public routes (read-only) - must be before wildcard routes
	r.GET("/posts", handlers.GetAllPosts)
	r.GET("/posts/:id/comments", handlers.GetCommentsByPost)
	r.GET("/posts/:id", handlers.GetPost)

	// Protected routes
	protected := r.Group("")
	protected.Use(middleware.AuthRequired())
	{
		// Posts
		protected.POST("/posts", handlers.CreatePost)
		protected.PUT("/posts/:id", handlers.UpdatePost)
		protected.DELETE("/posts/:id", handlers.DeletePost)

		// Comments
		protected.POST("/posts/:id/comments", handlers.CreateComment)
		protected.DELETE("/comments/:id", handlers.DeleteComment)
	}

	return r
}
