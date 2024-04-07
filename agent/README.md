# Agent

## Generate API Code

```zsh
go run github.com/ogen-go/ogen/cmd/ogen@latest --target ogen --clean api/openapi.json
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
