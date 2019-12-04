FROM golang:latest

WORKDIR /app

ADD . source/

RUN cd source && go build -o ../myapp

RUN rm -rf source/

ENTRYPOINT ["./myapp"]
