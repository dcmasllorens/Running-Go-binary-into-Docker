FROM ubuntu

RUN mkdir ./logs/
RUN chmod 777 ./logs/

RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates

COPY api_get /go/bin/

WORKDIR /go/bin/

CMD ["/go/bin/api_get"]