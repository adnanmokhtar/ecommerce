package http

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// Define a route group with a prefix
	countries := r.Group("/countries")

	// Register a GET route for listing countries
	countries.GET("/", controllers.list_countries_controller)
}
