# goproject
This is a demonstration of a simple API written in Go. In this example XML data is sent via a cURL request to the server. The data is processed (a specified field is made upper case) concurrently and returned as JSON.

## Installation & Usage
- Ensure Go is installed. Visit the [Go site](https://golang.org/doc/install) for instructions and documentation.

- Download and unzip file to your workspace directory, eg. $HOME/go/src.  For this example, we'll rename the folder to main.
- Navigate to the application folder. Build and run the *main.go* file from the command prompt eg.
  ```
  $ cd $HOME/go/src/main
  $ go build main.go
  $ ./main 
  ```
 - Send a cURL request to the application at localhost:8080/main
 ```
 curl -X POST -H 'Content-Type: application/xml' -d "<orderList>	
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
</orderList>" localhost:8080/main
```


## Another Test Case
- Example command string for another test:
```
curl -X POST -H 'Content-Type: application/xml' -d "<orderList>	
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
	
	
</orderList>" localhost:8080/main
```

