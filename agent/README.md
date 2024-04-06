# Agent

## Generate API Code

```bash
go run github.com/ogen-go/ogen/cmd/ogen@latest --target ogen --clean api/openapi.json
```

## Run Server

```bash
go run main.go handler.go agent.go backend.go llm.go
```

## Run Test

```bash
go test -v
```

```bash
% curl -X "POST" -H "Content-Type: application/json" http://localhost:8080/v0.0.1/chat
{"code":0,"message":"Bad Request"}%
```
