package routes

import "github.com/gin-gonic/gin"

func TestRoutes(r *Router) {
	r.engine.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "welcome test mode"})
	})
}