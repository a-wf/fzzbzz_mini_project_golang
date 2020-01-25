package controllers

import (
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/akdj/fizzbuzz/internal/services"
	"github.com/akdj/fizzbuzz/pkg/logger"
)


type Response struct {
	Answer	string `json:"answer"`
}

type Query struct {
	int1    int 
	int2 	int
	limit 	int 
	str1 	string 
	str2 	string 
}

func GetFizzBuzz(c *gin.Context) {

	var query Query

	if c.Bind(&query) == nil {
		int1, err1 := strconv.Atoi(c.Query("int1"))
		int2, err2 := strconv.Atoi(c.Query("int2"))
		limit, err3 := strconv.Atoi(c.Query("limit"))
		str1 := c.Query("str1")
		str2 := c.Query("str2")
		logger.Info("int1=%d; int2=%d; limit=%d; str1=%s; str2=%s", int1, int2, limit, str1, str2)

		if (err1 != nil || err2 != nil || err3 != nil || str2=="" || str1 == "" ) {
			logger.Error("Bad Request: bad input(s) type %s %s %s", err1, err2, err3);
			c.JSON(http.StatusBadRequest, "Bad Request: bad input(s) type")
			return
		}

		result := services.FizzBuzzAnswer(int1, int2, limit, str1, str2);

		resStruct := Response{ Answer: result}
	  	res, _ := json.Marshal(resStruct)
		logger.Info( string(res))
	
	  	c.JSON(http.StatusOK, resStruct)
	}
}

	