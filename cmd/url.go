package cmd

import (
	"douban/controller"
	"douban/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func URL() {
	router := gin.Default()
	router.Use(middleware.Cors()) //跨域请求问题解决

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //注册gin-swagger路由
	{
		router.GET("/wrongToken", middleware.JWTAuthMiddleware(), controller.GetErrToken) //返回浏览器错误token
		router.POST("/token", controller.GetToken)                                        //对外提供获取token的通道
		router.GET("type", middleware.JWTAuthMiddleware(), controller.GetMvsOfT)          //获取某种规则（url拼接参数）下的分类电影数据
		router.GET("rank", middleware.JWTAuthMiddleware(), controller.GetMvsOfR)          //获取某种规则（url拼接参数）下的排行榜电影数据(评分靠前的多个电影)
		router.GET("search", middleware.JWTAuthMiddleware(), controller.GetMvsOPerfs)     //搜索获取影视数据和影人数据
	}

	//user`s url
	usGroup := router.Group("/user")
	usGroup.Use(middleware.Cors())
	{
		usGroup.POST("loginByPwd", controller.LoginByPwd)                                           //密码登录
		usGroup.POST("registerOrLoginByPhone", controller.Register)                                 //注册\登录
		usGroup.POST("introduction", middleware.JWTAuthMiddleware(), controller.CreateIntroduction) //自我介绍
		usGroup.POST("sign", middleware.JWTAuthMiddleware(), controller.CreateSign)                 //签名
		usGroup.POST("avatar", middleware.JWTAuthMiddleware(), controller.UploadUserAvatar)         //上传用户头像api
		usGroup.GET("info", middleware.JWTAuthMiddleware(), controller.GetInfo)                     //获取自己主页的个人信息
		usGroup.GET("oInfo", middleware.JWTAuthMiddleware(), controller.GetOInfo)                   //获取其他人员的主页信息
		usGroup.GET("randCode", controller.GetRandCode)                                             //获取网站随机码
		usGroup.PATCH("alterPwd", middleware.JWTAuthMiddleware(), controller.UpdatePwd)             //更新密码
		usGroup.GET("question", middleware.JWTAuthMiddleware(), controller.GetQuestion)             //问题
		usGroup.GET("answer", middleware.JWTAuthMiddleware(), controller.GetAnswer)                 //答案
		usGroup.GET("movie/:label", middleware.JWTAuthMiddleware(), controller.GetWODMvs)           //获取当前用户想看和看过的电影
		usGroup.GET("lComment", middleware.JWTAuthMiddleware(), controller.GetLComments)            //获取用户发表的所有影评
	}
	//movie`s url
	mvGroup := router.Group("/movie")
	mvGroup.Use(middleware.Cors())
	{
		mvGroup.Use(middleware.JWTAuthMiddleware())
		mvGroup.GET(":mv_id/info", controller.GetMvInfo)                       //获取单个影视数据
		mvGroup.GET("hot", controller.GetHotMvs)                               //获取正在热映影视数据
		mvGroup.GET("future", controller.GetFutureMvs)                         //获取即将上映的电影数据
		mvGroup.POST(":mv_id/discuss", controller.CreateDiscuss)               //创建顶级讨论
		mvGroup.POST(":mv_id/:discuss_id/uDiscuss", controller.CreateUDiscuss) //创建顶级讨论下的回复
		mvGroup.DELETE(":mv_id/:discuss_id/discuss", controller.DelDisOfMv)    //删除电影下的顶级讨论及其所有回复
		mvGroup.PATCH(":mv_id/:discuss_id/uDiscuss", controller.DelUDisOfMv)   //删除电影顶级讨论下的回复(内容标记为【已删除】)
		mvGroup.GET(":mv_id/discuss", controller.GetDisOfMv)                   //获取电影的顶级讨论区
		mvGroup.GET(":mv_id/:discuss_id/uDiscuss", controller.GetUDisOfMv)     //获取电影讨论下的所有讨论
		mvGroup.GET(":mv_id/staff", controller.GetStaffsOfMv)                  //获取影视下的演职员数据
	}

	//comment`s url
	cGroup := router.Group("comment")
	cGroup.Use(middleware.Cors(), middleware.JWTAuthMiddleware())
	{
		cGroup.POST(":mv_id/sComment", controller.CreateWODMv)                  //添加当前用户想看或看过的电影，并添加可选数据：短评内容、短评标签、评分数据
		cGroup.DELETE(":mv_id/:comment_id/sComment", controller.DelScOfMv)      //删除当前短评，并删除想看和看过的电影数据和电影评分
		cGroup.GET(":mv_id/sComment", controller.GetScsOfMv)                    //获取电影下辖所有短评
		cGroup.POST(":mv_id/lComment", controller.CreateLComment)               //创建影评
		cGroup.POST(":mv_id/:comment_id/uLComment", controller.CreateULComment) //创建影评下的回应
		cGroup.DELETE(":mv_id/:comment_id/lComment", controller.DelLsOfMv)      //删除当前影评及回复
		cGroup.PATCH(":mv_id/:comment_id/uLComment", controller.DelULOfMv)      //删除当前影评下的回应(内容标记为【已删除】)
		cGroup.GET(":mv_id/lComment", controller.GetLcsOfMv)                    //获取当前电影所有影评
		cGroup.GET(":mv_id/:comment_id/uLComment", controller.GetUlsOfMv)       //获取当前影评下的所有回复
	}

	//staff`s url
	staffGroup := router.Group("/staff")
	staffGroup.Use(middleware.Cors(), middleware.JWTAuthMiddleware())
	{
		staffGroup.GET(":staff_id/info", controller.GetStaffInfo)   //获取具体电影数据
		staffGroup.GET(":staff_id/movie", controller.GetMvsOfStaff) //获取演职员下所参演的电影数据
	}
	router.Run(":8084")
}
