package main

import (
	"ProJectTest/BuildingAPI/Config"
	"ProJectTest/BuildingAPI/GeneraToken"
	"ProJectTest/BuildingAPI/Jwt"
	"github.com/gin-gonic/gin"
)

func main() {
	r:= gin.Default()
	r.POST("/Insert Into",Config.Create)
	r.GET("/Select",Config.Select)
	r.GET("/Find/:Id",Config.Find)
	r.PUT("/Update/:Id",Config.Update)
	r.DELETE("/Delete/:Id",Config.Delete)
	r.POST("/Login", Jwt.Login)
	r.POST("/LoginHash",Config.CreateHash)
	r.POST("/token",Jwt.LoginHash)
	api:=r.Group("secerity").Use(GeneraToken.AuTh())
	{
		r.POST("/api/select",Config.Select)
		r.POST("/api/generate",Jwt.LoginHash)
		api.GET("/api/ping",Config.Select)

	}
	r.Run()
}
