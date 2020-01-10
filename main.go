package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// Represents a <orderlist> element
type OrderList struct {
	XMLName xml.Name `xml:"orderList" json:"-"`
	Order   []Order  `xml:"order" json:"order"`
}

// Represents an <order> element
type Order struct {
	Id        string `xml:"id" json:"id"`
	Data      string `xml:"data" json:"data"`
	CreatedAt string `xml:"createdAt" json:"createdAt"`
	UpdatedAt string `xml:"updatedAt" json:"updatedAt"`
}

//Get and process data
func serveJSON(z http.ResponseWriter, req *http.Request) {

	//get data from request body
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	//log.Println(string(body))

	//create variable to hold XML data
	var data = &OrderList{}

	//Unmarshal Data
	err = xml.Unmarshal(body, data)
	if err != nil {
		panic(err)
	}

	//Create Waitgroup
	var wg sync.WaitGroup
	//Add wait group length based on number of orders
	wg.Add(len(data.Order))

	//For each Order
	for i := 0; i < len(data.Order); i++ {

		// create goroutine for processing Data item
		go func(i int) {
			data.Order[i].Data = strings.ToUpper(data.Order[i].Data)
			defer wg.Done()
		}(i)

	} //End For Loop
	wg.Wait() //Wait for goroutines to finish

	//Marshal data and output as JSON
	jsonData, _ := json.Marshal(data)
	fmt.Fprint(z, string(jsonData))
}

func main() {

	//serve JSON Data to server at port:8080
	http.HandleFunc("/", serveJSON)
	http.ListenAndServe(":8080", nil)

}
