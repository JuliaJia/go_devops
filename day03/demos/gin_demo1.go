package demos

import "github.com/gin-gonic/gin"

func Demo1() {
	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则
	r.GET("/")
}
