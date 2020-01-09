# goproject
This is a demonstration of a simple HTTP server written in Go. In this example XML data is parsed and processed (a specified field is  made upper case) concurrently and returned as JSON.

## Installation & Usage
- Ensure Go is installed. Visit the [Go site](https://golang.org/doc/install) for instructions and documentation.

- Unzip file to your workspace directory, eg. $HOME/go.
- Run the *main.go* file eg.
  ```
  $ cd $HOME/go/src/main
  $ ./main 
  ```
 - Open a browser and go to: [http://localhost:8080/main"](http://localhost:8080/main)

## Change Test Cases
- Open the *main.go* file
- On line 27, toggle between 'a' or 'b' as the assigned testXML string variable. The default is set to 'a'.

```
var testXML string = a  //This can be set to a or b
```



