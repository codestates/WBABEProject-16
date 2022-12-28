package router

import (
	"fmt"

	"codestates_lecture/WBABEProject-16/admincontroller"
	"codestates_lecture/WBABEProject-16/docs"
	"codestates_lecture/WBABEProject-16/logger"
	"codestates_lecture/WBABEProject-16/ordercontroller"

	"github.com/gin-gonic/gin"
	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"
)

type Router struct {
	adminController *admincontroller.Controller
	orderController *ordercontroller.Controller
}
func NewRouter(adminCtl *admincontroller.Controller,orderCtl *ordercontroller.Controller) (*Router, error){
	r := &Router{adminController: adminCtl, orderController: orderCtl}
	return r, nil
}
func liteAuth() gin.HandlerFunc {
	fmt.Println("auth 통과")
	return func(c *gin.Context) {
//~ 생략
		c.Next()
	}
}


//cross domain을 위해 사용
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//허용할 header 타입에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//허용할 method에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
func (r *Router)Idx() *gin.Engine {
	
	docs.SwaggerInfo.Host = "localhost:8080" //swagger 정보 등록
	docs.SwaggerInfo.Title= "pizza API"
    server := gin.New()
	server.Use(logger.GinLogger())
	server.Use(logger.GinRecovery(true))
	server.Use(CORS())
	url := ginSwg.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definitionv
	server.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler,url)) 

	server.GET("health", r.adminController.HealthCheck)
	admin := server.Group("admin")
	{
	  admin.POST("/category",r.adminController.AddCategory)
	  admin.PUT("/category",r.adminController.UpdateCategory)
	  admin.DELETE("/category",r.adminController.DeleteCategory)
	  admin.POST("/order/update", r.adminController.UpdateOrderStatus)
	} 
	order := server.Group("pizza")
	{
	   order.POST("/order", r.orderController.OrderPizza)
	   order.GET("/order/id/:id",r.orderController.GetOrderInfoById)
	   order.GET("/order/:name/:phone",r.orderController.FindOrderByNameAndPhone)
	}
	return server
 }