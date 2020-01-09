package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
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

var testXML string = a

// Test case data 1
var a string = `
<orderList>	
	<order>
		<id>aeffb38f-a1a0-48e7-b7a8-2621a2678534</id>
		<data>first_Order_Data</data>
		<createdAt>0001-01-01T00:00:00Z</createdAt>
		<updatedAt>0001-01-01T00:00:00Z</updatedAt>
	</order>
	<order>
		<id>beffb38f-b1a0-58e7-c7a8-3621a2678534</id>
		<data>second_Order_Data</data>
		<createdAt>0001-01-01T00:00:00Z</createdAt>
		<updatedAt>0001-01-01T00:00:00Z</updatedAt>
	</order>
</orderList>
`

//Test case data 2
var b string = `
<orderList>	
	<order>
		<id>asfd678f-df88-dfd8-dfd7-vc78bc87xv89</id>
		<data>order data 1</data>
		<createdAt>0001-01-01T00:00:00Z</createdAt>
		<updatedAt>0001-01-01T00:00:00Z</updatedAt>
	</order>
	<order>
		<id>asdf234-sd7f8-df8-asd98-vzx87vzx65asd</id>
		<data>order data 2</data>
		<createdAt>0001-01-01T00:00:00Z</createdAt>
		<updatedAt>0001-01-01T00:00:00Z</updatedAt>
	</order>
	<order>
		<id>asdf234-sd7f8-df8-asd98-vzx87vzx65asd</id>
		<data>order data 3</data>
		<createdAt>0001-01-01T00:00:00Z</createdAt>
		<updatedAt>0001-01-01T00:00:00Z</updatedAt>
	</order>
	<order>
		<id>asdf234-sd7f8-df8-asd98-vzx87vzx65asd</id>
		<data>order data 4</data>
		<createdAt>0001-01-01T00:00:00Z</createdAt>
		<updatedAt>0001-01-01T00:00:00Z</updatedAt>
	</order>
	
	
</orderList>
`

//create variable to hold XML data
var data = &OrderList{}

func main() {

	//Unmarshal Data
	err := xml.Unmarshal([]byte(testXML), data)
	if err != nil {
		log.Fatal(err)
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

	//serve JSON Data to server at port:8080
	http.HandleFunc("/", serveJSON)
	http.ListenAndServe(":8080", nil)

}

//JSON server function
func serveJSON(z http.ResponseWriter, r *http.Request) {
	//Marshal data and output as JSON
	jsonData, _ := json.Marshal(data)
	fmt.Fprint(z, string(jsonData))
}
