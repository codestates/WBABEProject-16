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


/*
피자라는 API를 따로 구현하기 보다는
api/v1/orders POST 의 방식을 이용해 orders라는 곳에서 한번에 처리할 수 있을 것 같습니다.
주문에 대한 메뉴로서 피자가 들어가게 되구요.
*/
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
	/*
	리소스가 올바르게 생성되면 일반적으로는 201 created를 return 합니다.
	*/
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

/*
이름과 전화번호를 통해서 주문을 가져오기 보다는, 
Order의 unique한 OrderId와 같은 값을 이용해서 해당 order를 가져오고, 그에 대한 정보를 가져오는 방식이 일반적입니다.
지금처럼 구성한다면, 필드가 늘어날 때마다 관련된 정보를 통해서 가져올 수 있는 API를 계속 생성해야 할 것입니다.
*/
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