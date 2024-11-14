package main

import(
    "log"
     "os"
    "QueueIt/routes"
    database "QueueIt/database"
	 "github.com/gin-gonic/gin"
    models "QueueIt/models"
)

func main() {
   port := os.Getenv("PORT")
   
   if port == "" {
	 port = "3000"
   }

   database.InitDB()
  
    
     err := database.DB.AutoMigrate(&models.User{})
     if err != nil {
         log.Fatalf("Failed to migrate models: %v", err)
     }
     log.Println("Database migration completed successfully.")

   router := gin.New()
   router.Use(gin.Logger())

   routes.AuthRoutes(router)
   routes.UserRoutes(router)

   router.GET("/api-1", func(c *gin.Context){
      c.JSON(200, gin.H{"success":"Access granted for api-1"})
   })

   router.GET("api-2", func(c *gin.Context){
      c.JSON(200, gin.H{"success":"Access granted for api-2"})
   })

   router.Run(":" + port)
}