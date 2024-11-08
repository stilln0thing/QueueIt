package database

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)


var DB *gorm.DB


func InitDB() {
 
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }


    postgresURL := os.Getenv("POSTGRES_URL")
    if postgresURL == "" {
        log.Fatal("POSTGRES_URL not found in environment")
    }

   
    db, err := gorm.Open(postgres.Open(postgresURL), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    // Configure the database connection pool
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Error configuring database connection pool: %v", err)
    }

    sqlDB.SetMaxIdleConns(10)               // Maximum idle connections
    sqlDB.SetMaxOpenConns(100)              // Maximum open connections
    sqlDB.SetConnMaxLifetime(5 * time.Minute) // Lifetime of a connection

    fmt.Println("Connected to PostgreSQL using GORM!")
    DB = db
}
