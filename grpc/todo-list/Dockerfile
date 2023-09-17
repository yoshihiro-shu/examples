FROM golang:1.21.1

WORKDIR /app
COPY . .

RUN go mod init todo
RUN go get -u google.golang.org/grpc
RUN go get -u google.golang.org/protobuf/cmd/protoc-gen-go
RUN go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative todo/todo.proto

EXPOSE 50051

CMD ["go", "run", "main.go"]
