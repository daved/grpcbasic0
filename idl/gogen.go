package idl

// alias
//go:generate -command pb-go protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:.
//go:generate -command pb-gg protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:.
//go:generate -command pb-sw protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:.

// generate
//go:generate pb-go ./grpcbasic0.proto
//go:generate pb-gg ./grpcbasic0.proto
//go:generate pb-sw ./grpcbasic0.proto

// add swagger def to package
//go:generate go-bindata -nocompress -pkg=idl ./grpcbasic0.swagger.json
