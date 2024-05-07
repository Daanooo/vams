package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	gin *gin.Engine
}

func NewRouter(db *gorm.DB) *Router {
	g := gin.Default()

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	// Register all custom route groups here
	registerDestinationRoutes(g)

	return &Router{
		gin: g,
	}
}

func (r *Router) Run() {
	r.gin.Run()
}
