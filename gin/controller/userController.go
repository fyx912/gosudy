package controller

import (
	"gostudy/gin/service"
	"log"
	// "encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginResponseBody struct {
	Username string `form:"user" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	// code string `form:"code" json:"code" binding:"required"`
}

func GetLoginHandler(this *gin.Context) {
	// this.Ctx.WriteString(" hello world")
	// this.HTML(http.StatusOK, "login.html",gin.H)
	// this.JSON(http.StatusOK, gin.H{
	// 	"message": "login",
	// })
	this.Header("Content-type", "text/html;charset=utf-8")
	this.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}
func PostLogin(this * gin.Context){
	var reqInfo LoginResponseBody
	if err := this.BindJSON(&reqInfo); err == nil {
		log.Println(reqInfo)
		if  "" != reqInfo.Username  {
			if reqInfo.Password != "" {
				if service.IsLogin(reqInfo.Username,reqInfo.Password) {
					log.Printf("username: %s ,password: %s", reqInfo.Username,reqInfo.Password)
					this.JSON(http.StatusOK, gin.H{"code":0,"msg":"success"})
				}else{
					this.JSON(http.StatusOK, gin.H{"code":0,"msg":"Failed! becesu: Incorrect user or password!"})
				}
			}else {
				log.Printf("Failed! becesu: Incorrect user password %s \n", reqInfo.Password)
				this.JSON(200, gin.H{"code":1,"msg":"Failed! becesu: Incorrect password!"})
			}
		}else{
			log.Printf("Failed! becesu: Incorrect user name %s \n", reqInfo.Password)
			this.JSON(200, gin.H{"code":1,"msg":"Failed! becesu: Incorrect user name!"})
		}
	}else{
		log.Printf("Failed! becesu:  %s \n", err.Error())
		this.JSON(200, gin.H{"code":2,"msg":"Failed! becesu : "+err.Error()})
	}
}

func GetUser(this *gin.Context){
	 jsonMap := make(map[string]interface{})
	jsonMap["code"] = 0
	jsonMap["msg"] = "success"
	user ,err := service.FindUserAll()
	if err != nil {
		log.Println(err,user)
		jsonMap["data"]= nil
	}else{
		jsonMap["data"]= user
	}
	this.JSON(http.StatusOK,jsonMap)
}	