# To build with arm architecture

* CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags '-s -w' -tags lambda.norpc -o ./bootstrap  ./main.go
* zip function.zip bootstrap