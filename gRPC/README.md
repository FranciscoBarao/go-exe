# gRPC doc
To generate gRPC code using .proto

```
export PATH="$PATH:$(go env GOPATH)/bin"
  
protoc --go_out=. --go-grpc_out=. folderPath/file.proto
```

How to test gRPC endpoints:
1. Create and use a Client
2. Use grpcurl -> Curl for gRPC



## Practical Examples

Generate gRPC code using .proto in proto/ folder
```
protoc --go_out=. --go-grpc_out=. proto/bookshop.proto
```

Basic request
```
grpcurl -plaintext localhost:9000 Inventory/GetBookList
```

Server side streaming
```
grpcurl -d '{"title": "testbook", "author":"testauthor", "page_count": 130, "language":"pt"}' -plaintext localhost:9000 Inventory/GetBookFieldsStream
```


Client side streaming. source: (gprcurl client side stream)[https://github.com/fullstorydev/grpcurl/issues/79]
```
grpcurl -plaintext -d @ localhost:9000 Inventory/CreateBookStream << EOF
{"key": "title", "value":"titleTest"}
{"key": "author", "value":"authorTest"}
EOF
```


Bidirectional streaming. Receives books and streams their fields back
```
grpcurl -plaintext -d @ localhost:9000 Inventory/BiGetBooksToFields << EOF
{"title": "1", "author":"1", "page_count": 1, "language":"1"}
{"title": "2", "author":"2", "page_count": 2, "language":"2"}
{"title": "3", "author":"3", "page_count": 3, "language":"3"}
EOF
```