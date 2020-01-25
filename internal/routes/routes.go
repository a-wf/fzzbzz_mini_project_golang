package routes

import (
	"github.com/akdj/fizzbuzz/internal/controllers"
	"github.com/gin-gonic/gin"
)


func Init() *gin.Engine {
	
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("fizzbuzz", controllers.GetFizzBuzz)
	}

	return r
}