An example for using buf.build/kei2100/protoc-gen-marshal-zap

buf.yaml
```yaml
# buf.yaml
version: v1
deps:
  - buf.build/kei2100/protoc-gen-marshal-zap
breaking:
  use:
    - FILE
lint:
  use:
    - DEFAULT
```

buf.gen.yaml
```yaml
# buf.gen.yaml
version: v1
plugins:
  - plugin: buf.build/protocolbuffers/go:v1.28.1
    out: gen/go
    opt: paths=source_relative
  - plugin: marshal-zap
    out: gen/go
    opt: paths=source_relative
```

generate *.pb.go and *.pb.marshal-zap.go
```bash
$ buf generate
```
