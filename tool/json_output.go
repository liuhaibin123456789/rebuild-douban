package tool

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type JsonFormat1 struct {
	Status int         `json:"status"`
	Err    string      `json:"err"`
	Data   interface{} `json:"data"`
}

type JsonFormat2 struct {
	Status int    `json:"status"`
	Err    string `json:"err"`
}

// JsonOutput1 data数据含有多个资源项，追加切片再放进来
func JsonOutput1(code string, err error, data interface{}, c *gin.Context) {
	c.JSON(200, gin.H{
		"status": code,
		"err":    fmt.Sprintf("%s", err),
		"data":   data,
	})
}

//JsonOutput2 没有资源项数据返回,帮助信息输出
func JsonOutput2(code string, err error, c *gin.Context) {
	c.JSON(200, gin.H{
		"status": code,
		"err":    fmt.Sprintf("%s", err),
	})
}
