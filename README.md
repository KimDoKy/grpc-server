# gRPC server

### directory structure

```bash
.
├── cmd        : 
├── config     : 환경변수들을 관리
├── gRPC
│   ├── client : gRPC server에 직접적인 call을 전송하는 client
│   ├── paseto : token 인증 방식 중 하나
│   ├── proto
│   └── server : gRPC server
├── go.mod
├── main.go
├── network    : router
├── repository : DB
├── service    : repository와 router를 연결
└── types      : 전역으로 사용되는 type을 관리
```

### install package

```bash
// for config file with toml
go get github.com/naoina/toml
// open source run server
go get github.com/gin-gonic/gin
// paseto (token auth)
go get github.com/o1egl/paseto
```

### install grpc, proto

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
$ export PATH="$PATH:$(go env GOPATH)/bin"
```
- [gRPC](https://grpc.io/docs/languages/go/quickstart/)

##### on Mac

```
$ brew install protobuf
```

#### Regenerate gRPC code

```
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    gRPC/proto/auth.proto
```

after
```
gRPC/proto
├── auth.pb.go
├── auth.proto
└── auth_grpc.pb.go
```

### Run

```bash
(project root)$ go run .
```
