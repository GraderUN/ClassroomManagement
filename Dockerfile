FROM golang

ENV GO111MODULE=on

WORKDIR /go/src/github.com/GraderUN/ClassroomManagement

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 8080
ENTRYPOINT ["/go/src/github.com/GraderUN/ClassroomManagement"]