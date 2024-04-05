# Backend

## Generate API Code

```bash
go run github.com/ogen-go/ogen/cmd/ogen@latest --target ogen --clean api/openapi.json
```

## Run Server

```bash
go run main.go handler.go logic.go
```

## Run Test

```bash
go test -v
```

```bash
% curl -X "POST" -H "Content-Type: application/json" http://localhost:8080/v0.0.1/talk
{"code":0,"message":"Bad Request"}%
```
