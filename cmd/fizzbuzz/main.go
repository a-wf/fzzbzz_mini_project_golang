package main

import (
	"io"
	"os"
	"net/http"
	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"
	"github.com/akdj/fizzbuzz/pkg/logger"
	"github.com/akdj/fizzbuzz/internal/routes"
)

func Init() {
	// logger.Init("debug", "/var/log")
	viper.SetConfigName("fizzbuzz_config")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	logger.CheckIfError(err, "ReadInConfig")

	gin.DisableConsoleColor()
	f, _ := os.Create(viper.GetString("log.file.path")+viper.GetString("log.file.name")+".log")
	gin.DefaultWriter = io.MultiWriter(f)

	logger.Info("Init");
}


func main() {
	Init()
	router := routes.Init()
	var err error
	
	if(viper.GetString("server.protocol")== "http"){
		err = http.ListenAndServe(":"+ viper.GetString("server.port"), router)

	}else if(viper.GetString("server.protocol")== "https"){
		err = http.ListenAndServeTLS(":"+ viper.GetString("server.port"), 
			viper.GetString("server.ssl.certificate"),  viper.GetString("server.ssl.key"), router)
	}
	if err != nil {
			logger.Error("ListenAndServe: %s", err)
	}
}