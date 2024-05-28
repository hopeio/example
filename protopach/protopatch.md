https://github.com/alta/protopatch
go get github.com/alta/protopatch
go install github.com/alta/protopatch/cmd/protoc-gen-go-patch@latest
cd protopatch
之前的命令：
protoc -I. --go_out=paths=source_relative:. *.proto
tag注入
protoc -I. -Ipatch --go-patch_out=plugin=go,paths=source_relative:.  *.proto