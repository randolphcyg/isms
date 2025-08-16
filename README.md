# ISMS

# 更新文档
```
make api
```

## proto 生成 pb service
```
kratos proto client api/isms/v1/os.proto
kratos proto server api/isms/v1/os.proto

kratos proto client api/isms/v1/industry.proto
kratos proto server api/isms/v1/industry.proto

kratos proto client api/isms/v1/dashboard.proto
kratos proto server api/isms/v1/dashboard.proto

# software
protoc --proto_path=./api \
       --proto_path=./third_party \
       --go_out=paths=source_relative:./api \
       --go-http_out=paths=source_relative:./api \
       --go-grpc_out=paths=source_relative:./api \
       --validate_out=paths=source_relative,lang=go:./api \
       api/isms/v1/software.proto
```

## Docker 部署
```bash
# build
docker build -t isms:1.0 .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf isms:1.0
```