# urlShortener
Important to say, this is a learning project to understand GoLang principles and practice.

## How to use it
Download the project and start the server
```shell script
go run main.go
```
Perform a POST request
```shell script
curl --location --request POST 'http://localhost:8080/encode' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url": "http://www.google.com.br"
}'
```
return must be something like
```shell script
{
    "redirect_url": "http://localhost:8080/go/xcjHoW17"
}
```
now you can access your local browser
http://localhost:8080/go/xcjHoW17

## Improvements
- Statistics for how many stored urls and redirects performed