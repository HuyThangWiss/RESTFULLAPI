package Function

import (
	"ProJectTest/BuildingAPI/InformationAPI"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Res struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
}

func ReadAll(cin *gin.Context)  {
	resp,err := http.Get("http://localhost:8080/Select")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	cin.JSON(http.StatusCreated,gin.H{"Data  find":map[string]interface{}{"Data":string(body)}})

}
func ReadAllToken(cin *gin.Context)  {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", "http://localhost:8080/secerity/api/ping", nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Cell", `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6NCwiTmFtZSI6IiQyYSQxMCRhdXlFMnZPSXk1R1ZVUnE2bHdLWnBPVUJnWWFCdDFJdFlLMVN2ellUbUJCQ1ZNdTlRT0VkeSIsImV4cCI6MTY2NzQ5MDk4N30.Ap2CQXOJloYCEIcET8SaMZ_PNhFD676Ckck9bRbIIE0`)
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
	var arr[] InformationAPI.Books
	json.Unmarshal(body, &arr)
	//myJson, _ :=json.Marshal(body)
	cin.JSON(http.StatusOK,gin.H{"Data ":Res{
		Code: 0,
		Data: string(body),
	}})
}
