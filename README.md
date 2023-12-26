# Distributed_Anomaly_Detection_System

The implementation of anomaly detection with Golang Deep Learning in distributed setting up

## Features:
- one master node, several slave nodes
- communicate through TCP protocols
- messages are json packets
- each node has its unique ID
- first node is default the master cluster. If it dies, re-elect a master

## Structure

- cmd: contains the main appliations and entry points

- internal: defined packages with internal usage only

- pkg: packages used for external projects

- api: define api-related code

- scripts: automation scripts

- web: web-related assets, static files and templates

## References

*https://dev.to/tikazyq/golang-in-action-how-to-implement-a-simple-distributed-system-2n0n*
*https://didil.medium.com/building-a-simple-distributed-system-with-go-consul-39b08ffc5d2c*
*https://confusedcoders.com/general-programming/go-lang/create-a-basic-distributed-system-in-go-lang-part-1*
*https://confusedcoders.com/general-programming/go-lang/create-a-basic-distributed-system-in-go-lang-part-2-http-server-json-requestresponse*
