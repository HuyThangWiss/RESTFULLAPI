package main

import (
	"ProJectTest/Server/Function"
	"github.com/gin-gonic/gin"
)


func main() {
	r:=gin.Default()
	r.GET("/ServerA/Select",Function.ReadAllToken)
	r.Run(":8081")
}
