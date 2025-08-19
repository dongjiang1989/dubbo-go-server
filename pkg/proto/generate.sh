export GO111MODULE="on"
export GOPROXY="http://goproxy.io"
go install github.com/dubbogo/tools/cmd/protoc-gen-go-triple@v1.0.9
protoc --go_out=. --go-triple_out=. node.proto