FROM golang:latest
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
VOLUME /go/src/app/
RUN go mod download
EXPOSE 8080
CMD "go" "run" "main.go"