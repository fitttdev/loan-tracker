package main

import (
    "github.com/gin-gonic/gin"
    "loan-tracker/internal/api/routes"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
    r := gin.Default()
    
    // Load HTML templates
    
    routers.SetupHomeRouter(r) // Set up home-related routes
    routers.SetupLoanRouter(r) // Set up loan-related routes

    return r
}

func main() {
    r := setupRouter()
    r.Run(":8080")
}
