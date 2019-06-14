package main

import (
    "github.com/micro/go-log"
	"net/http"
	"iHomeMicroMicro/homeweb/handler"
    "github.com/micro/go-web"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// create new web service
    service := web.NewService(
            web.Name("go.micro.web.homeweb"),
            web.Version("latest"),
            web.Address(":8888"),
    )

	// initialise service
    if err := service.Init(); err != nil {
            log.Fatal(err)
    }
	rou := httprouter.New()
	//映射静态页面
	rou.NotFound = http.FileServer(http.Dir("html"))

	//获取地区信息
	rou.GET("/api/v1.0/areas",handler.GetArea)
	//获取图片验证码
	rou.GET("/api/v1.0/imagecode/:uuid",handler.GetImageCd)
	//获取短信验证码
	rou.GET("/api/v1.0/smscode/:mobile",handler.Getsmscd)
	//获取session
	rou.GET("/api/v1.0/session",handler.GetSession)
	//注册
	rou.POST("/api/v1.0/users",handler.Postreg)
	//登陆
	rou.POST("/api/v1.0/sessions",handler.PostLogin)
	//推出登陆
	rou.DELETE("/api/v1.0/session", handler.DeleteSession)
	//请求用户基本信息 GET /api/v1.0/user
	rou.GET("/api/v1.0/user", handler.GetUserInfo)

	////上传头像 POST
	//rou.POST("/api/v1.0/user/avatar",handler.Postavatar)
	////请求更新用户名 PUT
	//rou.PUT("/api/v1.0/user/name",handler.Putuserinfo)
	////实名认证检查 GET
	//rou.GET("/api/v1.0/user/auth",handler.Getuserauth)
	////实名认证 post
	//rou.POST("/api/v1.0/user/auth",handler.Postuserauth)


	////请求当前用户已发布房源信息  GET
	//rou.GET("/api/v1.0/user/houses",handler.Getuserhouses)
	////发布房源信息 POST
	//rou.POST("/api/v1.0/houses",handler.Posthouses)
	////上传房源图片信息  POST
	//rou.POST("/api/v1.0/houses/:id/images",handler.Posthousesimage)
	////请求房源详细信息 GET
	//rou.GET("/api/v1.0/houses/:id",handler.Gethouseinfo)
	////首页轮播图请求  Get  Index
	//rou.GET("/api/v1.0/house/index",handler.GetIndex)
	//// 搜索  api/v1.0/houses?aid=5&sd=2017-11-12&ed=2017-11-30&sk=new&p=1
	//rou.GET("/api/v1.0/houses",handler.Gethouses)
	//
	////post 发布订单 api/v1.0/orders
	//rou.POST("/api/v1.0/orders",handler.Postorders)
	////get 查看房东/租客订单信息请求
	//rou.GET("/api/v1.0/user/orders",handler.Getuserorder)
	////put房东同意/拒绝订单
	////api/v1.0/orders/:id/status
	//rou.PUT("/api/v1.0/orders/:id/status",handler.Putorders)
	////PUT 用户评价订单信请求
	////api/v1.0/orders/:id/comment
	//rou.PUT("/api/v1.0/orders/:id/comment",handler.Putcomment)



	// register html handler
	service.Handle("/", rou)

	// run service
    if err := service.Run(); err != nil {
            log.Fatal(err)
    }
}
