package controller

import (
	"codestates_lecture/WBABEProject-16/model"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Controller struct {
	mongoDb *model.Model
}
type DeleteRequestBody struct{
	Name string `json:name "binding":"required"`
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
// @Param model.PizzaCategory body model.PizzaCategory true "PizzaCategory Info"
// @Router /admin/category [post]s
// @Success 200 {object} model.PizzaCategory
func (ctl *Controller) AddCategory(c *gin.Context){
   var requestBody model.PizzaCategory
   if err := c.ShouldBind(&requestBody); err != nil {
	   fmt.Println(err)
	   c.JSON(http.StatusBadRequest, gin.H{"result":false})
	   return
   }
   ctl.mongoDb.AddCategory(requestBody)
   c.JSON(http.StatusOK, gin.H{"result":true})
}



// Post
// @Summary update a pizza in category
// @Description 피자정보를 update하는 api
// @Accept  json
// @Produce  json
// @Param model.PizzaCategory body model.PizzaCategory true "PizzaCategory Info"
// @Router /admin/category [put]s
// @Success 200 {object} model.PizzaCategory
func (ctl *Controller) UpdateCategory(c *gin.Context){
	var requestBody model.PizzaCategory
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
	fmt.Print(requestBody)
	result, err := ctl.mongoDb.DeleteByName(requestBody.Name)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"result":errors.New("error")})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result":result})
 }
 


// GetOk godoc
// @Summary health check
// @Description health check
// @Router /health [get]s
func (ct *Controller) HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "health")
}