package controller

import (
	"awesomeProject1/commen"
	"awesomeProject1/dto"
	"awesomeProject1/model"
	"awesomeProject1/response"
	"awesomeProject1/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	db := commen.GetDb()
	fmt.Println(db)
	var ReU = model.User{}
	err := c.Bind(&ReU)
	if err != nil {
		return
	}
	name := ReU.Name
	iphone := ReU.Iphone
	password := ReU.Password

	if len(iphone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号长度不合适")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码长度不合适")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(8)
	}
	if util.IsIphone(db, iphone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}

	hasedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	user := model.User{
		Name:     name,
		Iphone:   iphone,
		Password: string(hasedpassword),
	}
	db.Create(&user)
	token, err := commen.ReleaseToken(ReU)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "token创建失败"})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"data":    gin.H{"token": token},
		"massage": "注册成功",
	})

}

func Login(c *gin.Context) {
	db := commen.GetDb()
	fmt.Println(db)
	user := model.User{}
	iphone := c.PostForm("iphone")
	password := c.PostForm("password")

	if len(iphone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": "手机号长度不合适"})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": "密码长度不合适"})
		return
	}

	if !util.IsIphone(db, iphone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": "用户不存在"})
		return
	} else {
		if user = util.IfPassword(db, iphone, password); user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
			return
		}
	}
	token, err := commen.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "token创建失败"})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"data":    gin.H{"token": token},
		"massage": "注册成功",
	})
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.GetUserDto(user.(model.User))}})
}
