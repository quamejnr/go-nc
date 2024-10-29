## What is this?
This is a poor man implementation of netcat in Go. It doesn't even have any flags

## Why this?
Built this during my experimentation with computer networking.

## How can you use this?
1. Clone the repo locally using `git clone https://github.com/quamejnr/go-nc.git`
2. Run `main.go` in the root directory and and pass in the address you want to connect to. You can also build the code using `go build`
```Golang
go run main.go example.com:http
```
3. This will connect you to `example.com` you can then send http requests over like you would do in netcat
```sh
GET / HTTP/1.1
HOST: example.com
```
4. Press enter twice and it should return an http response.

