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


func NewController(mongo *model.Model) (*Controller, error){
	r := &Controller{mongoDb: mongo}
	return r, nil
}

type nameRequestBody struct {
	Name string
}

type pnumRequestBody struct {
	Pnum string
}

type updateRequestBody struct {
	Name nameRequestBody
	Pnum pnumRequestBody
}



type personRequestBody struct{
	Name string `bson:string`
	Age string  `bson:string`
	Pnum string  `bson:string`
}



// PostOK godoc
// @Summary get person Info, return peronInfo by json.
// @Description api test를 위한 기능.
// @name personByName
// @Accept  json
// @Produce  json
// @Param nameRequestBody body nameRequestBody true "user name"
// @Router /common/personByName [post]s
// @Success 200 {object} personRequestBody
func (ct *Controller) GetPersonByName(c *gin.Context){
	
    var requestBody nameRequestBody;
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println("error")
	}
	fmt.Println(requestBody.Name)
	result, err := ct.mongoDb.GetPersonByName(requestBody.Name)
	if err != nil {
		fmt.Println("error 발생")
	}
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{"data":result})
}

func (ct *Controller) GetPersonByPnum(c *gin.Context){
	
    var requestBody pnumRequestBody;
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println("error")
	}
	fmt.Println(requestBody.Pnum)
	result, err := ct.mongoDb.GetPersonByPnum(requestBody.Pnum)
	if err != nil {
		fmt.Println("error 발생")
	}
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{"data":result})
}


// PostOK godoc
// @Summary join a person , return objectID
// @Description api test를 위한 기능.
// @name Join
// @Accept  json
// @Produce  json
// @Param model.Person body model.Person true "person Info"
// @Router /acc/v01/Join [post]s
func (ct *Controller) JoinPerson(c *gin.Context){
	var requestBody personRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println("err")
	}
	result, err := ct.mongoDb.JoinPerson(model.Person(requestBody))
	c.JSON(http.StatusOK, gin.H{"data":"success"})
	if err != nil {
		fmt.Println("error 발생")
	}
	c.JSON(http.StatusOK, gin.H{"objectsId":result})
}


// PostOK godoc
// @Summary update person age , return update_count 
// @Description api test를 위한 기능.
// @name updatePerson
// @Accept  json
// @Produce  json
// @Param model.Person body model.Person true "person Info"
// @Router /acc/v01/updatePerson [put]s
func (ct *Controller) UpdatePerson(c *gin.Context){
	var requestBody model.Person
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println("err")
	}
	result, err := ct.mongoDb.UpdatePerson(requestBody)
	if err != nil {
		fmt.Println("error")
	}

	c.JSON(http.StatusOK, gin.H{"update_count":result})
}

func (ct *Controller) DeletePerson(c *gin.Context){
	var requestBody model.Person
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println("err")
	}
	result, err := ct.mongoDb.DeletePerson(requestBody)
	if err != nil {
		fmt.Println("error")
	}

	c.JSON(http.StatusOK, gin.H{"delete_count":result})
}





// GetOk godoc
// @Summary health check
// @Description health check
// @Router /health [get]s
func (ct *Controller) HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "health")
}