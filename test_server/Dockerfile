FROM golang:1.21.3@sha256:24a09375a6216764a3eda6a25490a88ac178b5fcb9511d59d0da5ebf9e496474
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd cmd
COPY pkg pkg
RUN go build -o bin/server ./cmd/server/...
EXPOSE 8888
ENTRYPOINT [ "/app/bin/server" ]
