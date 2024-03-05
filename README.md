# Distributed_Anomaly_Detection_System

The implementation of anomaly detection with Golang Deep Learning in distributed setting up

## Features:
- one master node, several slave nodes
- communicate between components via gRPC protocols
- messages are json packets
- each node has its unique ID
- first node is default the master cluster. If it dies, re-elect a master
- central node is to aggregate the results from client decision maker
- data query service with GraphQL or Mongodb or Redis

## Structure

- cmd: contains the main appliations and entry points

- internal: defined packages with internal usage only

    - dataparser: process multiple data (logs, network traffic, process)

    - detection: design the detection logic for different data sources

    - aggregator: the aggregation process with centrol node

- pkg: packages used for external projects

- api: define api-related code

- scripts: automation scripts

- web: web-related assets, static files and templates

## Workflow

- Client initializes API call for Master Node
- Master node assigns task to work node n
- Work node n executes script/task and report status to Master Node
- MongoDB + GraphQL for data storage and query support

## Running Instructions (Mac)

```
# install gRPC
go install google.golang.org/protobuf/cmd/protoc-gen-go@1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
export PATH="$PATH:$(go env GOPATH)/bin"

# download protobuf3 (need to add option with go_package in .proto file )
brew install protobuf
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
protoc --go_out=. {name}.proto
```

## References

*https://dev.to/tikazyq/golang-in-action-how-to-implement-a-simple-distributed-system-2n0n*
*https://confusedcoders.com/general-programming/go-lang/create-a-basic-distributed-system-in-go-lang-part-1*
*https://confusedcoders.com/general-programming/go-lang/create-a-basic-distributed-system-in-go-lang-part-2-http-server-json-requestresponse*
*https://www.apollographql.com/blog/using-graphql-with-golang*
