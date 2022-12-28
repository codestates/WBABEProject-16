package admincontroller

import (
	"codestates_lecture/WBABEProject-16/model"
	"codestates_lecture/WBABEProject-16/structs"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Controller struct {
	mongoDb *model.Model
}
type DeleteRequestBody struct{
	PizzaId string `binding:"required"`
}

type OrderStatusBody struct {
	Id string `binding:"required"`
	Status string `binding:"required"`
}

func NewController(mongo *model.Model) (*Controller, error){
	r := &Controller{mongoDb: mongo}
	return r, nil
}



// Post
// @Summary add a pizza in category
// @Description 피자종류를 추가하는 api
// @Accept  json
// @Produce  json
// @Param   structs.RequestPizzaCategoryBody body  structs.RequestPizzaCategoryBody true "PizzaCategory Info"
// @Router /admin/category [post]s
// @Success 200 {object} model.PizzaCategory
func (ctl *Controller) AddCategory(c *gin.Context){
   var requestBody structs.RequestPizzaCategoryBody
   if err := c.ShouldBind(&requestBody); err != nil {
	   fmt.Println(err,"err")
	   c.JSON(http.StatusBadRequest, gin.H{"result":false})
	   return
   }
  result, err := ctl.mongoDb.AddCategory(requestBody)
   if err != nil {
	c.JSON(http.StatusOK, gin.H{"result":result})
	return
   }
   c.JSON(http.StatusOK, gin.H{"result":result})
   
}



// Post
// @Summary update a pizza in category
// @Description 피자정보를 update하는 api
// @Accept  json
// @Produce  json
// @Param  structs.RequestPizzaCategoryBody body structs.RequestPizzaCategoryBody true "PizzaCategory Info"
// @Router /admin/category [put]s
// @Success 200 {object} model.PizzaCategory
func (ctl *Controller) UpdateCategory(c *gin.Context){
	var requestBody  structs.RequestPizzaCategoryBody
	if err := c.ShouldBind(&requestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"result":false})
		return
	}
	result, err := ctl.mongoDb.UpdateCategory(requestBody)
	if err != nil{
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"result":errors.New("error")})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result":result})
 }
 

 // Delete
// @Summary delete a pizza in category
// @Description 피자정보를 update하는 api
// @Accept  json
// @Produce  json
// @Param DeleteRequestBody body DeleteRequestBody true "delete"
// @Router /admin/category [delete]s
func (ctl *Controller) DeleteCategory(c *gin.Context){

	var requestBody DeleteRequestBody
	if err := c.ShouldBind(&requestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"result":false})
		return
	}
	result, err := ctl.mongoDb.DeleteByName(requestBody.PizzaId)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"result":errors.New("error")})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result":result})
 }
 



// statusUpdate godoc
// @Summary update a status in order
// @Description 피자주문정보에서 주문접수, 조리, 배달완료 등 상태를 update하는 API
// @Accept  json
// @Produce  json
// @Param OrderStatusBody body OrderStatusBody true "update a status in order"
// @Router /admin/order/update [post]s
 func(ctl *Controller) UpdateOrderStatus(c *gin.Context){
    var requestBody OrderStatusBody
	if err := c.ShouldBind(&requestBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"result":false})
		return
	}

	result, resultErr := ctl.mongoDb.UpdateOrderStatus(requestBody.Id, requestBody.Status)
	if resultErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result":false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"update_result":result})
 }


// GetOk godoc
// @Summary health check
// @Description health check
// @Router /health [get]s
func (ct *Controller) HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "health")
}