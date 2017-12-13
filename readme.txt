export GOPATH=$(pwd) 或 (export GOPATH=.../grpc)
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}    	   //需要到GOPATH中执行
go get -u google.golang.org/grpc                             	   //工作目录中执行(src中)
protoc --go_out=plugins=grpc:./输出目录／ ./*.proto文件所在目录    //将$GOPATH/bin添加到 PATH执行才不会报错（自动生成服务器和客户端的注册函数）
(如果用$ protoc --go_out=./go/ ./proto/helloworld.proto生成*.pb.go需要自己写服务器和客户端的注册函数）

将生成的proto/*.pb.go文件拷贝到相应的文件夹里(./calc)

go run server/server.go
go run client/client.go
