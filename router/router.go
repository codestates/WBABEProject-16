package router

import (
	"fmt"

	controller "codestates_lecture/WBABEProject-16/controller"
	"codestates_lecture/WBABEProject-16/docs"
	"codestates_lecture/WBABEProject-16/logger"

	"github.com/gin-gonic/gin"
	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"
)


type Router struct {
	controller *controller.Controller
}
func NewRouter(ctl *controller.Controller) (*Router, error){
	r := &Router{controller: ctl}
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
	docs.SwaggerInfo.Title= "hello my api"
    server := gin.New()
	server.Use(logger.GinLogger())
	server.Use(logger.GinRecovery(true))
	server.Use(CORS())
	url := ginSwg.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definitionv
	server.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler,url)) 

	server.GET("health", r.controller.HealthCheck)
	account := server.Group("common")
	{
	
		account.POST("/personByName",liteAuth(),r.controller.GetPersonByName)	
		account.POST("/personByPnum",liteAuth(),r.controller.GetPersonByPnum)	
		account.POST("/join", liteAuth(),r.controller.JoinPerson)
        account.PUT("/updatePerson", liteAuth(), r.controller.UpdatePerson)
		account.DELETE("/deletePerson",liteAuth(), r.controller.DeletePerson)

		// account.GET("/getPerson", p.ct.GetPerson)	// controller 패키지의 실제 처리 함수
		// account.POST("/updatePerson",p.ct.UpdatePerson)
	}

	return server
 }