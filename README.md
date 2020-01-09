# goproject
This is a demonstration of a simple HTTP server written in Go. In this example XML data is parsed and processed (a specified field is  made upper case) concurrently and returned as JSON.

## Installation & Usage
- Ensure Go is installed. Visit the [Go site](https://golang.org/doc/install) for instructions and documentation.

- Unzip file to your workspace directory, eg. $HOME/go.
- Run the *main.go* file from the command prompt eg.
  ```
  $ cd $HOME/go/src/main
  $ ./main 
  ```
 - Open a browser and go to: [http://localhost:8080/main](http://localhost:8080/main)
 - The following JSON output should appear:
 ```
 {"order":[{"id":"aeffb38f-a1a0-48e7-b7a8-2621a2678534","data":"FIRST_ORDER_DATA","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":"beffb38f-b1a0-58e7-c7a8-3621a2678534","data":"SECOND_ORDER_DATA","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}]}
 ```

## Change Test Cases
- Open the *main.go* file
- On line 27, toggle between 'a' or 'b' as the assigned testXML string variable (the default is set to 'a'). Save the *main.go* file.

```
var testXML string = a  //This can be set to a or b
```
- Rebuild the application from the application directory, eg. $HOME/go/src/main
```
$ go build main.go
```
- Refresh the browser to see the newly assigned data.



