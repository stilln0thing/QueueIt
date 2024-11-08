package main

import(
    "QueueIt/routes"
	 "os"
	 "github.com/gin-gonic/gin"
)

func main() {
   port := os.Getenv("PORT")

   if port == "" {
	 port = "3000"
   }

   router := gin.New()
   router.Use(gin.Logger())

   routes.AuthRoutes(router)
}