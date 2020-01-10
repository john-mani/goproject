# goproject
This is a demonstration of a simple HTTP server written in Go. In this example XML data is parsed and processed (a specified field is  made upper case) concurrently and returned as JSON.

## Installation & Usage
- Ensure Go is installed. Visit the [Go site](https://golang.org/doc/install) for instructions and documentation.

- Download and unzip file to your workspace directory, eg. $HOME/go/src.  For this example, we'll rename the folder to main.
- Navigate to the application folder. Build and run the *main.go* file from the command prompt eg.
  ```
  $ cd $HOME/go/src/main
  $ go build main.go
  $ ./main 
  ```
 - Send a cURL request to localhost:8080
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


## Change Test Cases
- Open the *main.go* file
- On line 27, toggle between 'a' or 'b' as the assigned testXML string variable (the default is set to 'a'). Save the *main.go* file.

```
var testXML string = a  //This can be set to a or b
```
- Rebuild and run the application from the application directory, eg. $HOME/go/src/main
```
$ go build main.go
$ ./main
```
- Refresh the browser to see the newly assigned data.



