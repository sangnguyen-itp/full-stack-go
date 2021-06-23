package routes

import "full-stack-go/controllers"

func DevRoutes(r *Router) {
	r.engine.GET("/", controllers.IndexPage())
}
