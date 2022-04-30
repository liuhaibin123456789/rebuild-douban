package controller

import (
	"douban/dao"
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// LoginByPwd
// @Summary 登录
// @accept mpfd
// @Produce  json
// @Param phone formData string true "手机号"
// @Param password formData string true "密码"
// @Success 200 {object} string "成功"
// @Header  200  {string}  Token     "token"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /user/loginByPwd [post]
func LoginByPwd(c *gin.Context) {
	phone := c.PostForm("phone")
	pwd := c.PostForm("password")

	token, err := service.LoginByPwd(phone, pwd)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
			"token":  "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
		"token":  token,
	})
}

// Register
// @Summary 注册用户，如果已经注册过，则会切换至登录
// @accept mpfd
// @Produce  json
// @Param user formData model.User true "用户信息"
// @Success 200 {object} string    "成功"
// @Header 200  {string}  Token     "token"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /user/registerOrLoginByPhone [post]
func Register(c *gin.Context) {
	var u model.User
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  err,
		})
		return
	}
	token, err := service.Register(u)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf(err.Error()),
			"token":  "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
		"token":  token,
	})
}

//CreateIntroduction 创建用户介绍
func CreateIntroduction(c *gin.Context) {
	phone := c.GetString("phone")
	if phone == "" {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  "failed to get the 'phone' in the gin context,please check the JWTMiddleware in the application",
		})
		return
	}
	userIntroduction := c.PostForm("user_introduction")
	err := service.CreateIntroduction(phone, userIntroduction)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
	})
}

//CreateSign 创建用户签名
// @summary  获取token
// @accept mpfd
// @produce  json
// @Param    user_sign formData  string  true  "用户签名"  maxlength(200)
// @success  200       {object}  tool.JsonFormat1  "成功"
// @failure  200       {object}  tool.JsonFormat1  "请求错误"
func CreateSign(c *gin.Context) {
	phone := c.GetString("phone")
	if phone == "" {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  "failed to get the 'phone' in the gin context,please check the JWTMiddleware in the application",
		})
		return
	}
	sign := c.PostForm("user_sign")
	err := service.CreateSign(phone, sign)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
	})
}

//UploadUserAvatar 上传用户头像
func UploadUserAvatar(c *gin.Context) {
	//从上下文中获取用户标识
	phone := c.GetString("phone")
	if phone == "" {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  "failed to get the 'phone' in the gin context,please check the JWTMiddleware in the application",
		})
		return
	}
	//获取头像文件
	fileHeader, err := c.FormFile("user_avatar")
	if err != nil {
		c.JSON(500, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
		})
		return
	}
	//拼接相对路径
	dst := fmt.Sprintf(".././douban/static/picture/useravatar/%s", fileHeader.Filename) //todo 抽离至配置文件
	//保存文件
	err = c.SaveUploadedFile(fileHeader, dst)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
		})
		return
	}
	//保存图像名
	err = service.UploadUserAvatar(phone, fileHeader.Filename)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
	})
}

//GetInfo 获取用户部分信息
func GetInfo(c *gin.Context) {
	//从上下文中获取用户标识
	phone, exists := c.Get("phone")
	if !exists {
		c.JSON(200, gin.H{
			"status":    "0",
			"error":     "failed to get the 'phone' in the gin context,please check the JWTMiddleware in the application",
			"user_info": "",
		})
		return
	}
	us, err := service.GetInfo(phone.(string))
	if err != nil {
		c.JSON(200, gin.H{
			"status":    "0",
			"error":     fmt.Sprintf("%s", err),
			"user_info": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":    "1",
		"error":     "",
		"user_info": us,
	})
}

func GetOInfo(c *gin.Context) {
	phone := c.PostForm("phone")
	oInfo, err := service.GetOInfo(phone)
	if err != nil {
		c.JSON(200, gin.H{
			"status":    "0",
			"error":     fmt.Sprintf("%s", err),
			"user_info": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":    "1",
		"error":     "",
		"user_info": oInfo,
	})
}

//GetRandCode 获取网站随机码
// @Summary 获取网站随机验证码
// @accept mpfd
// @Produce  json
// @Param phone query string true "手机号" minlength(11) maxlength(11)
// @Success 200 {object} int "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /user/randCode [get]
func GetRandCode(c *gin.Context) {
	phone := c.Query("phone")
	randCode, err := service.GetRandCode(phone)
	if err != nil {
		c.JSON(200, gin.H{
			"status":    "0",
			"error":     fmt.Sprintf("%s", err),
			"rand_code": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":    "1",
		"error":     "",
		"rand_code": randCode,
	})
}

func UpdatePwd(c *gin.Context) {
	phone := c.PostForm("phone")
	oldPassword := c.PostForm("old_password")
	newPassword := c.PostForm("new_password")

	err := service.UpdatePwd(phone, oldPassword, newPassword)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
	})
}

func GetAnswer(c *gin.Context) {
	phone, res := c.Get("phone")
	if !res {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  "user_name is not in the context",
			"answer": "",
		})
		return
	}
	answer, err := service.GetAnswer(phone.(string))
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
			"answer": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
		"answer": answer,
	})

}

func GetQuestion(c *gin.Context) {
	phone, res := c.Get("phone")
	if !res {
		c.JSON(200, gin.H{
			"status":   "1",
			"error":    "the phone is not in the gin of context,please check status of login!",
			"question": "",
		})
		return
	}
	question, err := service.GetQuestion(phone.(string))
	if err != nil {
		c.JSON(200, gin.H{
			"status":   "0",
			"error":    fmt.Sprintf("%s", err),
			"question": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":   "1",
		"error":    "",
		"question": question,
	})
}

// GetToken godoc
// @summary  获取token
// @accept mpfd
// @produce  json
// @Param    phone     formData  string  true  "手机号,11位数字"  length(11)
// @param    password  formData  string  true  "密码,8到16位数字大小写字母组合"   minlength(8)  maxlength(16)
// @success  200       {object}  tool.JsonFormat1  "成功"
// @failure  200       {object}  tool.JsonFormat2  "请求错误"
// @router   /token [post]
func GetToken(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user model.UserInfo
	err := c.ShouldBind(&user)
	data := gin.H{"token": ""}
	if err != nil {
		tool.JsonOutput2("0", err, c)
		return
	}
	//校验用户名和密码是否正确
	rightPwd, err := dao.SelectUserPwd(user.Phone)
	if err == nil && rightPwd == tool.Encrypt(user.Password) {
		// 生成Token
		tokenString, err := tool.CreateToken(user.Phone)
		if err != nil {
			tool.JsonOutput2("0", err, c)
			return
		}
		data["token"] = tokenString
		tool.JsonOutput1("0", nil, data, c)
		return
	}
	tool.JsonOutput2("0", nil, c)
}

//GetErrToken 服务端返回错误token
func GetErrToken(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
		"token":  "this is a wrong token",
	})
}

// GetWODMvs
// @Summary 获取用户想看或看过的电影
// @Produce  json
// @Param label path string true "想看的标签为0,看过的标签为1" enum("1","0") default("1")
// @Success 200 {array} model.OfMovie "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /user/movie/{label} [get]
func GetWODMvs(c *gin.Context) {
	label := c.Param("label")
	phone := c.GetString("phone")
	wodMvs, err := service.GetWODMvs(label, phone)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
			"movies": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
		"movies": wodMvs,
	})
}

func GetLComments(c *gin.Context) {
	phone := c.GetString("phone")
	lComments, err := service.GetLComments(phone)
	if err != nil {
		c.JSON(200, gin.H{
			"status":        "0",
			"error":         fmt.Sprintf("%s", err),
			"long_comments": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":        "1",
		"error":         "",
		"long_comments": lComments,
	})
}
