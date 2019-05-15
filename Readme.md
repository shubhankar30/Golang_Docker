## Docker Important Notes

- To run go code within docker, an exec file needs to be executed.  
- The exec file can be created with many ways. 
1. During building of docker container 
2. Previously built using go build and then executed within docker container or passed during runtime 


- In this approach, main.go is built in docker file, and to execute it, docker compose is used 

```
docker build -t test .

(test is image name)
```

Run commands:
``` 
- Go build main.go     ---> Creates main exec file 
- Docker run –it –p 8080:8080 –t test ./main    ----> To run the exec file 
```

```
- Docker run –it –p 8080:8080 –t test go run go/src/temp/main.go   --> run the go file within docker command 

In this case, no need to add RUN go build go/src/temp/main.go in Dockerfile 
```

```
- Docker build –t test 
- Docker run –it –p 8080:8080 –t test 
```

```
docker-compose up -d --build
```

 