package controller

import (
	"codestates_lecture/WBABEProject-16/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Controller struct {
	mongoDb *model.Model
}

type PizzaOrderBody struct {
	PersonInfo model.OrderPersonInfo
	OrderInfo model.OrderInfo
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
// @Param PizzaOrderBody body PizzaOrderBody true "PizzaOrderBody Info"
// @Router /pizza/order [post]s
func (ctl *Controller) OrderPizza(c *gin.Context){
	var requestBody PizzaOrderBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		  fmt.Println("err")
	}
	result, resultErr := ctl.mongoDb.OrderPizza(requestBody.OrderInfo, requestBody.PersonInfo)
	if resultErr != nil {
          c.JSON(http.StatusBadRequest, gin.H{"result":false})
	}
	c.JSON(http.StatusOK, gin.H{"result":result})
	
}