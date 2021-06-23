package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexPage() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// Call the HTML method of the Context to render a template
		ctx.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)
	}
}
