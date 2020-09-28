FROM golang

ENV GO111MODULE=on

# Add Maintainer Info
LABEL maintainer="Sanhernandezmon <sanhernandezmon@unal.edu.co>"

WORKDIR /go/src/github.com/GraderUN/ClassroomManagement

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 8080
ENTRYPOINT ["/github.com/GraderUN/ClassroomManagement"]