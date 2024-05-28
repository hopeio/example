https://github.com/alta/protopatch

cd protopatch
之前的命令：
protoc -I. --go_out=paths=source_relative:. *.proto
tag注入
protoc -I. -Ipatch --go-patch_out=plugin=go,paths=source_relative:.  *.proto