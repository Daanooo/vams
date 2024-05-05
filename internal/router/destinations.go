package router

import "github.com/gin-gonic/gin"

// Name is configusing
func registerDestinationRoutes(g *gin.Engine) {
	routes := g.Group("destinations")

	// TODO remove this example
	routes.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{})
	})
}
