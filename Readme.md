To run go code within docker, an exec file needs to be executed.  

 

The exec file can be created with many ways. 

 

During building of docker container 

 
Or previously built using go build 

 And then executed within docker container  

Or the exec file can be passed during runtime 

In lessontime, main.go is build in docker file, and to execute it, docker compose is used 



1. 
- Go build main.go     ---> Creates main exec file 
- Docker run –it –p 8080:8080 –t test ./main    ----> To run the exec file 

2.
- Docker run –it –p 8080:8080 –t test go run go/src/temp/main.go   --> run the go file within docker command 

In this case, no need to add RUN go build go/src/temp/main.go in Dockerfile 

3.

- Docker build –t test 
- Docker run –it –p 8080:8080 –t test 

 

 