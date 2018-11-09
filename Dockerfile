FROM golang:1.8

ADD . go/src/temp

#RUN go build go/src/temp/main.go

EXPOSE 8080
