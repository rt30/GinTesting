package main


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)
type Order struct {
	OrderId int `json:"oderId"`
	CustomerName string `json:"custName"`
	OrderReview string `json:"review"`
}
// slice of orders

var orders = [] Order{
	Order{101,"rt","good"},
	Order{102,"naveen","not so good ambience"},
	Order{103,"dp","good food"},
	Order{104,"sal","nice arrangement"},
}
func PostOrder(c *gin.Context){
	body := c.Request.Body

	content, err := ioutil.ReadAll(body)

	if err != nil{
		fmt.Println("Sorry, no content found: ", err.Error())
	}

	fmt.Println(content)
	c.JSON(http.StatusCreated, gin.H{
		"message" :string(content),
	})
}
func GetOrders(c *gin.Context){
	c.JSON(200,&orders)
}

func HomePage(c *gin.Context){
	c.String(200,"Pong")
}

func Authenticate(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	if secret, ok := secrets[user]; ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}

func setupAPI() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	//authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
	//	"foo" : "bar",
	//}))
	//
	////authorized.GET("/secrets", Authenticate)
	//authorized.GET("/secrets", func(c *gin.Context) {
	//	// get user, it was set by the BasicAuth middleware
	//	user := c.MustGet(gin.AuthUserKey).(string)
	//	if secret, ok := secrets[user]; ok {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	//	}
	//})

	api.GET("/orders",GetOrders)
	api.GET("/", HomePage)
	api.POST("/PostOrder",PostOrder)
	return r
}


var secrets = gin.H{
	"foo":    gin.H{"name": "rt", "phone": "gin"},
}

func main() {
	r := setupAPI()
	r.Run(":8080")
}