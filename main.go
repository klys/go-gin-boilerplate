package main

import (

  "github.com/gin-gonic/gin"
  "gintest/routes"
)

func main() {
  r := gin.Default()
  routes.Test(r)
  r.Run("localhost:9090")
   // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}