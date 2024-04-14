# gRPC server

### directory structure

```bash
.
├── cmd        : 
├── config     : 환경변수들을 관리
├── gRPC
│   ├── client
│   ├── paseto
│   ├── proto
│   └── server
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
```

### Run

```bash
(project root)$ go run .
```
