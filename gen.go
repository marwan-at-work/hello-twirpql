package hellotwirpql

//go:generate protoc --go_out=. --twirp_out=. --twirp_swagger_out=. service.proto
//go:generate protoc --twirpql_out=. service.proto
