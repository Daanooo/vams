package router

import "github.com/gin-gonic/gin"

type Router struct {
	gin *gin.Engine
}

func NewRouter() *Router {
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
