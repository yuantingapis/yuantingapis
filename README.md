# Yuanting APIs

This repository contains the original interface definitions of public Yuanting APIs that support both REST and gRPC protocols.

## Compile proto files and generate Go clients

```bash
protoc -I . -I $API_COMMON_PROTOS \
    --go_out=../go-genproto \
    --go_opt=paths=source_relative \
    --go-grpc_out=../go-genproto \
    --go-grpc_opt=paths=source_relative \
    yuanting/yt/v1/*.proto
```

## Generate go client

```bash
protoc -I . -I $API_COMMON_PROTOS \
    --go_gapic_out ../go-client \
    --go_gapic_opt 'go-gapic-package=github.com/yuantingapis/go-client/yuanting/yt/v1;yt' \
    --go_gapic_opt 'module=github.com/yuantingapis/go-client' \
    yuanting/yt/v1/*.proto
```
