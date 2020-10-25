# urlShortener
Important to say, this is a learning project to understand GoLang principles and practice.

## How to use it
Download the project and start the server
```shell script
go run main.go
```
Perform a POST request
```shell script
$ curl --location --request POST 'http://localhost:8080/encode?url=http://www.google.com'
```
return must be something like
```shell script
/go/JJxrKUOX
```
now you can access your local browser
http://localhost:8080/go/JJxrKUOX

## Improvements
- Uses form instead of url param
- Statistics for how many stored urls and redirects performed
- Add test cases