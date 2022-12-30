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
		/*
		CORS 허용을 위해서 모든 도메인을 허용한다면 보안에 이슈가 생깁니다. 
		보통 운영되는 시스템의 경우는 특정한 도메인만을 허용하고 그 이외의 요청은 거부하도록 설정합니다.
		*/
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
	/*
	엔드포인트 구성에 대해서 전반적인 코멘트 드립니다.
	1. REST API의 성숙도 모델에 대해서 공부해보시면 좋을 것 같습니다.

	2. 일반적으로 HTTP URI에 update 와 같은 행위는 들어가지 않습니다. 
		복수형의 단어로 구성을 하고, 동일한 URI 내에서 http method만 변경하여 행위를 표현하는 것이 일반적인 REST API의 구성 방식입니다.

		e.g.
		GET v1/orders -> 주문 목록을 조회.
		GET v1/orders/1 -> 1번 주문 조회.
		POST v1/orders -> 주문 생성.
		PATCH v1/orders/1 -> 1번 주문에 대해서 업데이트
		DELETE v1/orders/1 -> 1번 주문에 대해서 삭제
	*/

	docs.SwaggerInfo.Host = "localhost:8080" //swagger 정보 등록
	docs.SwaggerInfo.Title= "pizza API"
    server := gin.New()
	server.Use(logger.GinLogger())
	server.Use(logger.GinRecovery(true))
	server.Use(CORS())
	url := ginSwg.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definitionv
	server.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler,url)) 

	server.GET("health", r.adminController.HealthCheck)
	/*
	그룹을 통해 라우팅을 나누어 주신 점 좋습니다. 코드를 읽기에 용이하네요.
	*/
	admin := server.Group("admin")
	{
	  admin.POST("/category",r.adminController.AddCategory)
	  admin.PUT("/category",r.adminController.UpdateCategory)
	  admin.DELETE("/category",r.adminController.DeleteCategory)
	  admin.POST("/order/update", r.adminController.UpdateOrderStatus)
	}
	/*
	주문과 관련된 내용인데 실제 그룹명은 Pizza인 이유가 있을까요?
	*/ 
	order := server.Group("pizza")
	{
	   order.POST("/order", r.orderController.OrderPizza)
		/*
		id가 두번 들어간 이유는 무엇인가요?
		*/
	   order.GET("/order/id/:id",r.orderController.GetOrderInfoById)
	   order.GET("/order/:name/:phone",r.orderController.FindOrderByNameAndPhone)
	}
	return server
 }