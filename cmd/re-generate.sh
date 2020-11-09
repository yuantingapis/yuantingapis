#!/usr/bin/env -S bash -e

source cmd/lib.sh || {
  echo "Are you at repo root?"
  exit 1
}

# API_COMMON_PROTOS refers https://github.com/googleapis/api-common-protos
# Generate Go protos.
runcmd protoc -I . -I $API_COMMON_PROTOS \
  --go_out=../go-genproto \
  --go_opt=paths=source_relative \
  --go-grpc_out=../go-genproto \
  --go-grpc_opt=paths=source_relative \
  yuanting/yt/v1/*.proto || {
  err
  exit 1
}

# Generate Go client.
runcmd protoc -I . -I $API_COMMON_PROTOS \
  --go_gapic_out=../yuanting-go \
  --go_gapic_opt 'go-gapic-package=github.com/yuantingapis/yuanting-go/yuanting/yt/v1;yt' \
  --go_gapic_opt 'module=github.com/yuantingapis/yuanting-go' \
  --go_gapic_opt="grpc-service-config=./yuanting/yt/v1/yt_grpc_service_config.json" \
  yuanting/yt/v1/*.proto || {
  err
  exit 1
}

# Generate OpenAPI document.
runcmd protoc -I . -I $API_COMMON_PROTOS \
  --grpc-gateway_out ../go-genproto \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt paths=source_relative \
  yuanting/yt/v1/*.proto || {
  err
  exit 1
}
