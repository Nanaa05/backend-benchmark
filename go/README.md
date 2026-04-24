```
GOMAXPROCS=$(nproc) go run main.go

go build -o server
GOMAXPROCS=$(nproc) ./server
```
