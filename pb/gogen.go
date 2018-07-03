package pb

// alias
//go:generate -command pb-go protoc -I../proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:.
//go:generate -command pb-gg protoc -I../proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:.
//go:generate -command pb-sw protoc -I../proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:.

// generate
//go:generate pb-go ../proto/grpcbasic0.proto
//go:generate pb-gg ../proto/grpcbasic0.proto
//go:generate pb-sw ../proto/grpcbasic0.proto

// add swagger def to package
//go:generate go-bindata -nocompress -pkg=pb ./grpcbasic0.swagger.json
