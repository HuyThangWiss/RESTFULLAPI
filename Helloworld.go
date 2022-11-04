package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var learns []Learn

type Code struct {
	Code int
	Body []Learn
}
type Code1 struct {
	Code int
	Body ID
}
type Learn struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Duration   int    `json:"duration"`
	Created_by string `json:"created_by"`
}
type ID struct {
	ID int
}

func main() {

	handleRequests()
}
func handleRequests() {
	http.HandleFunc("/", SimpleServer)
	http.HandleFunc("/api/tasks", getListTask)
	http.HandleFunc("/api/long-tasks", getLongTasks)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func ReadFile() []byte {
	jsonFile, err := os.Open("todo.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}
func getLongTasks(w http.ResponseWriter, request *http.Request) {
	json.Unmarshal(ReadFile(), &learns)

	b := make([]Learn, 0)
	json.Unmarshal(ReadFile(), &learns)
	for i := 0; i < len(learns); i++ {
		if learns[i].Duration >= 8 {
			b = append(b, learns[i])
		}
	}

	Code := Code{0, b}
	newb, _ := json.MarshalIndent(Code, "", " ")
	fmt.Fprintf(w, string(newb))
}
func SimpleServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
func getListTask(w http.ResponseWriter, r *http.Request) {
	json.Unmarshal(ReadFile(), &learns)
	switch r.Method {
	case "GET":
		Code := Code{0, learns}
		newb, _ := json.MarshalIndent(Code, "", " ")
		fmt.Fprintf(w, string(newb))
	case "POST":
		var newtask Learn
		err := json.NewDecoder(r.Body).Decode(&newtask)
		fmt.Println(err)
		if err != nil {
			fmt.Print(err)
		}
		newtask.ID = learns[len(learns)-1].ID + 1
		//b, _ := json.MarshalIndent(newtask, "", "  ")

		c := make([]Learn, 0)
		for i := 0; i < len(learns); i++ {
			c = append(c, learns[i])
		}
		c = append(c, newtask)
		file, _ := json.MarshalIndent(c, "", " ")
		_ = ioutil.WriteFile("todo.json", file, 0644)
		ID := ID{newtask.ID}

		Code1 := Code1{0, ID}
		newb, _ := json.MarshalIndent(Code1, "", " ")
		fmt.Fprintf(w, string(newb))
	}

}
