FROM golang:1.8



ADD . /go/src/temp

RUN rm -rf /go/src/temp/static

RUN ls /go/src/temp

#RUN go build go/src/temp/main.go

EXPOSE 8080
