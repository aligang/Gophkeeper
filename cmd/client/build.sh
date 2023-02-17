go build -ldflags "-X github.com/aligang/Gophkeeper/internal/buildinfo.Version=v1.0.0 -X 'github.com/aligang/Gophkeeper/internal/buildinfo.BuildTime=$(date +'%Y/%m/%d %H:%M:%S')'" -o gophkeeper-cli client.go

GOOS=windows GOARCH=amd64 go build -ldflags "-X github.com/aligang/Gophkeeper/internal/buildinfo.Version=v1.0.0 -X 'github.com/aligang/Gophkeeper/internal/buildinfo.BuildTime=$(date +'%Y/%m/%d %H:%M:%S')'" -o gophkeeper.exe client.go