package main

import( 
	"os"
	"github.com/gin-gonic/gin"
	routes "golang-RESTAURANTS_MANAGEMENT_IN_GO/routes"
	// "golang-RESTAURANTS_MANAGEMENT_IN_GO/database"
	// "golang-RESTAURANTS_MANAGEMENT_IN_GO/middleware"
	// "go.mongodb.org/mongo-driver/mongo"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port= "8000";
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	routes.AuthRoutes(router)
	
}	