# Backend

## Generate API Code

```bash
go run github.com/ogen-go/ogen/cmd/ogen@latest --target ogen --clean openapi.json
```

## Run Server

```bash
go run main.go internal.go
```

## Test

```bash
% curl -X "POST" -H "Content-Type: application/json" --data "{\"prompt\":\"Cat\"}" http://localhost:8080/talk
{"visualizableData":"Cat"}%
```
