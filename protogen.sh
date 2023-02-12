export PATH=$PATH:$GOPATH/bin
protoc \
  --go_out=. \
  --go_opt=module=github.com/aligang/Gophkeeper \
  --go-grpc_out=. \
  --go-grpc_opt=module=github.com/aligang/Gophkeeper \
  proto/config/config.proto \
  proto/account/account.proto \
  proto/account/account_service.proto \
  proto/secret/secret.proto \
  proto/secret/secret_service.proto \
  proto/token/token.proto \
  proto/pipeline/tree.proto
