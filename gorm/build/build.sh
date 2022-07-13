echo building...
# export GOOS=linux
# export GOARCH=arm64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o ../cmd/test2  ../internal/app/main.go
echo success