package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ExitGet(c *gin.Context) {
	log.Println("@@@[ExitGet]Ready to log out.")
	//清除该用户登录状态的数据
	session := sessions.Default(c)
	session.Delete("loginuser")
	session.Save()
	//session.Clear()

	fmt.Println("delete session...", session.Get("loginuser"))
	c.Redirect(http.StatusMovedPermanently, "/")
}
