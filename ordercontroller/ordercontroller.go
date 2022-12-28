package ordercontroller

import (
	"codestates_lecture/WBABEProject-16/model"
	"codestates_lecture/WBABEProject-16/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Controller struct {
	mongoDb *model.Model
}

type RequestNameAndPhone struct {
	Name string `bind:"required"`
	Phone string `bind:"required"`
}



func NewController(mongo *model.Model) (*Controller, error){
	r := &Controller{mongoDb: mongo}
	return r, nil
}



// Post
// @Summary order a pizza 
// @Description 피자를 주문하는 API
// @Accept  json
// @Produce  json
// @Param structs.RequestPizzaOrderBody body structs.RequestPizzaOrderBody true "RequestPizzaOrderBody Info"
// @Router /pizza/order [post]s
func (ctl *Controller) OrderPizza(c *gin.Context){
	var requestBody structs.RequestPizzaOrderBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		  
		  fmt.Println(err,"err")
	}
	fmt.Println(requestBody)
	result, resultErr := ctl.mongoDb.OrderPizza(requestBody.OrderInfo,requestBody.PersonInfo)
	if resultErr != nil {
          c.JSON(http.StatusBadRequest, gin.H{"result":result})
	}
	c.JSON(http.StatusOK, gin.H{"order_id":result})
	
}



// Get
// @Summary 주문번호를 통해서 주문정보를 받아볼 수 있는 API
// @Description 주문번호를 통해서 주문정보를 받아볼 수 있는 API
// @Accept  json
// @Produce  json
// @Param id path string true "order id"
// @Router /pizza/order/id/{id} [get]
func (ctl *Controller) GetOrderInfoById(c *gin.Context){
	var objectId =  c.Param("id")
	result, err := ctl.mongoDb.FindOrderById(objectId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result":false})
	}
	c.JSON(http.StatusOK, gin.H{"data":result})
}


// Get
// @Summary 이름과 전화번호를 통해서 주문내역을 확인할 수 있는 API
// @Description 이름과 전화번호를 통해서 주문내역을 확인할 수 있는 API
// @Accept  json
// @Produce  json
// @Param name path string true "user name"
// @Param phone path string true "user phone"
// @Router /pizza/order/{name}/{phone} [get]
func (ctl *Controller) FindOrderByNameAndPhone(c *gin.Context){
	var name =  c.Param("name")
	var phone =  c.Param("phone")
	result := ctl.mongoDb.FindOrderByNameAndPhone(name, phone)
	c.JSON(http.StatusOK, gin.H{"data":result})
}