# Agent

## Generate API Code

```zsh
go run github.com/ogen-go/ogen/cmd/ogen@latest --target ogen --clean api/openapi.json
```

## Change generated code

Change `newServerConfig` function in `ogen/oas_cfg_gen.go` as follows:

```go
func newServerConfig(opts ...ServerOption) serverConfig {
  cfg := serverConfig{
    NotFound: http.NotFound,
    MethodNotAllowed: func(w http.ResponseWriter, r *http.Request, allowed string) {
      status := http.StatusMethodNotAllowed
      if r.Method == "OPTIONS" {
        w.Header().Set("Access-Control-Allow-Methods", allowed)
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")  // Add
        w.Header().Set("Access-Control-Max-Age", "3600")  // Add
        status = http.StatusNoContent
      } else {
        w.Header().Set("Allow", allowed)
      }
      w.WriteHeader(status)
    },
    ErrorHandler:       ogenerrors.DefaultErrorHandler,
    Middleware:         nil,
    MaxMultipartMemory: 32 << 20, // 32 MB
  }
  for _, opt := range opts {
    opt.applyServer(&cfg)
  }
  cfg.initOTEL()
  return cfg
}
```

## Create schema.sql

```zsh
$ brew install sqlite
$ sqlite3 chinook.db
sqlite> .tables
albums          employees       invoices        playlists
artists         genres          media_types     tracks
customers       invoice_items   playlist_track
sqlite> .schema  # Copy schema to schema.sql
sqlite> .exit
```

## Run Server

```zsh
go run main.go handler.go agent.go backend.go llm.go
```

## Run Test

```zsh
go clean -testcache
go test -v
```

```zsh
% curl -X "POST" -H "Content-Type: application/json" --data "{\"prompt\":\"Cat\"}" http://localhost:8080/v0.0.1/chat
{"visualizableData":"Cat"}
% curl -X "POST" -H "Content-Type: application/json" http://localhost:8080/v0.0.1/chat
{"code":0,"msg":"Bad Request"}
% curl -X "POST" -H "Content-Type: application/json" --data "{\"prompt\":\"Cat\"}" http://localhost:8080/v0.0.1/xxx
404 page not found
```
