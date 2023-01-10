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
	/*
	Des는 무슨 의미인가요? 기본적으로 축약형 보다는 풀어쓰는 것이 조금 더 직관적입니다.
	*/
	Des string `binding:"required"`
	M int `binding:"required"`
	L int `binding:"required"`
	Limit_Order int `binding:"required"`
	/*
	필드명만 보고서는 bool 형태인지 유추하기가 힘들어 보입니다. 
	*/
	Order_status bool `binding:"required"`
  }


