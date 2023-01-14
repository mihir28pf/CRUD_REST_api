FROM golang:latest
ENV GOPATH=/app
COPY . .
RUN go get
RUN go build -o main .
CMD [ "go", "run", "main.go" ]
