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
  
    
   modelsList := []interface{}{
    &models.User{},
    &models.Business{},
    &models.Event{},
}

// Iterate over the models and apply AutoMigrate
for _, model := range modelsList {
    err := database.DB.AutoMigrate(model)
    if err != nil {
        log.Fatalf("Failed to migrate model %v: %v", model, err)
    }
    log.Printf("Migration for model %v completed successfully.", model)
}

   router := gin.New()
   router.Use(gin.Logger())

   routes.AuthRoutes(router)
   routes.UserRoutes(router)
   routes.BusinessRoutes(router)

   // router.GET("/api-1", func(c *gin.Context){
   //    c.JSON(200, gin.H{"success":"Access granted for api-1"})
   // })

   // router.GET("api-2", func(c *gin.Context){
   //    c.JSON(200, gin.H{"success":"Access granted for api-2"})
   // })

   router.Run(":" + port)
}