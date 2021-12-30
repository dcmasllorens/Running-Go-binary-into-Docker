# Runing Go binary into Docker

1. We need to compile the .go main file, for compile we need to have  [installed](https://go.dev/doc/install) go in our computer.

~~~~
go build api_get.go
~~~~

2.  We will create a ***Dockerfile***, in this case we will use a docker image ubuntu, because if we use go docker image, we can't do http.get because this image d'ont have the library system for to do this. And the other reason is that we compile the go binary before.

~~~~
FROM ubuntu

RUN mkdir ./logs/
RUN chmod 777 ./logs/

RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates

COPY api_get /go/bin/

WORKDIR /go/bin/

CMD ["/go/bin/api_get"]
~~~~

3. When we have the Dockerfile, we will create the ***docker-compose.yaml***, this file will contain the configuration of directory logs and the name of the image. The directory log, will create in your current directory when you start the Docker image

~~~~
version: '3.3'
services:
  myapp:
    container_name: docker-api
    build: .
    image: docker-api
    volumes:
      - ./logs_api/:/go/bin/logs/
~~~~

4. In this case we will mount a Docker Volume, configured with docker-compose, because we want to see the log that API generates. 

First we need to configure the Docker Compose, to use docker compose we need to [install the docker Engine](https://docs.docker.com/engine/install/debian/), when have installed the docker Engine then we can [install the docker Compose](https://docs.docker.com/compose/install/) 

5. When we have all installed and configured we will build the image with command docker-compose, the parameter -f is for specifying the compose file. 

~~~~
docker-compose -f docker-compose.yml build
~~~~

With this command we create a docker image. And before we need to start the image. With the parameter -d we put the container in background.

~~~~
docker-compose up -d
~~~~

For know if container api-docker is running.

~~~~
docker ps
~~~~

This binary will execute when the container starts, when this happens this will do infinite requests every 4 seconds, and this will stop when the container stops.

## References

[Website](https://www.cloudbees.com/blog/building-minimal-docker-containers-for-go-applications) explain how to build docker container with go aplication 

[Docker compose documentation](https://docs.docker.com/compose/)

[Docker volumes documentation](https://docs.docker.com/storage/volumes/)