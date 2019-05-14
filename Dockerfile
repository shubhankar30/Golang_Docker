FROM golang:latest



ADD . /go/src/temp
# RUN go get -u "github.com/go-sql-driver/mysql"

RUN rm -rf /go/src/temp/static

RUN ls /go/src/temp

#RUN go build go/src/temp/main.go
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/gin-gonic/gin"

EXPOSE 8080
