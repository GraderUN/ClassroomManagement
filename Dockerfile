FROM golang:latest

ENV GO111MODULE=on

# Add Maintainer Info
LABEL maintainer="Sanhernandezmon <sanhernandezmon@unal.edu.co>"

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

WORKDIR /go/src/github.com/GraderUN/ClassroomManagement

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 8080
ENTRYPOINT ["/go/src/github.com/GradeUN/ClassroomManagement"]