package structs




type RequestOrderPersonInfo struct {
	Name string 
	Phone string
	Address string 
	PizzaId string 

  }
  
  type RequestOrderInfo struct {
	PizzaId string 
	Size string 
	Amount int 
	PersonId string 
	Status string 
  }
  type RequestPizzaOrderBody struct {
	  PersonInfo RequestOrderPersonInfo 
	  OrderInfo RequestOrderInfo
  }


  type RequestPizzaCategoryBody struct {
	Name string `binding:"required"`
	Des string `binding:"required"`
	M int `binding:"required"`
	L int `binding:"required"`
	Limit_Order int `binding:"required"`
	Order_status bool `binding:"required"`
  }


