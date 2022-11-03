package Controler

import (
	"awesomeProject/B3/InforJwtB3"
	"bytes"
	json "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func B44(cin *gin.Context)  {
	values := map[string]string{"foo": "baz"}
	jsonData, err := json.Marshal(values)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/Select", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	// appending to existing query args
	q := req.URL.Query()
	q.Add("User", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtYXN2IjoiMDA1IiwicGFzc3dvcmRzdiI6IjEyMzQ1NjciLCJleHAiOjE2NjM5Mzc3NjEsImlhdCI6MTY2MzY3ODU2MX0.D8hAZCzSUKF7Ztc7EPa8UssshkGaz56r9g7rthMkzvg")

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	cin.JSON(http.StatusOK,gin.H{"Data ":string(responseBody)})

}

type Res struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
}

func B55(cin *gin.Context)  {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", "http://localhost:8080/api/Select", nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("User", `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtYXN2IjoiMDA1IiwicGFzc3dvcmRzdiI6IjEyMzQ1NjciLCJleHAiOjE2NjM5Mzc3NjEsImlhdCI6MTY2MzY3ODU2MX0.D8hAZCzSUKF7Ztc7EPa8UssshkGaz56r9g7rthMkzvg`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err!= nil{
		return
	}
	var arr[] InforJwtB3.Sinhviens
	json.Unmarshal(body, &arr)
	//myJson, _ :=json.Marshal(body)
	cin.JSON(http.StatusOK,gin.H{"Data ":Res{
		Code: 0,
		Data: body,
	}})
}





//var (
//	listUser = make([]User, 0)
//)
//func (user User) ToString() string {
//	json.Unmarshal(ReadFile(), &user)

//
//func Get(c *gin.Context)  {
//	Name:=c.DefaultQuery("Name","Gusert")
//	var data = map[string]interface{}{
//		"Message ":"Hello from "+Name+" get ping",
//	}
//	c.JSON(http.StatusOK,data)
//}
//

//func GetId(c *gin.Context)  {
//
//	id:=c.Param("id")
//	c.JSON(http.StatusOK,gin.H{
//		"id ":id,
//	})
//}