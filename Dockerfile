FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go get -d -v ./...
RUN go install -v ./...

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o film-app ./cmd/main.go

CMD ["./film-app"]