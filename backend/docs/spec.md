# Specification

## Sequence Diagram

```mermaid
sequenceDiagram
  box User System
    actor User
    participant Frontend
    participant Controller
    participant User DB
  end
  box QueryTalk
    participant Backend
    participant LLM
    participant DB
  end

  User ->> Frontend: Input prompt
  Note over Frontend, Controller: REST API
  Frontend ->> Controller: Request
  Note over Controller, Backend: gRPC
  Controller ->> Backend: Request
  Backend ->> DB: Check permission
  DB ->> Backend: Return result
  alt Unauthorized
    Backend ->> Controller: 401 Error
    Controller ->> Frontend: 401 Error
    Frontend ->> User: 401 Error
  end
  Backend ->> LLM: Send prompt
  LLM ->> Backend: Return SQL query
  Backend ->> Controller: Return SQL query
  Controller ->> User DB: Execute SQL query
  User DB ->> Controller: Return result
  Controller ->> Frontend: Return visualizable data
  Frontend ->> User: Visualize data
```

## Technologies

### Frontend

- TypeScript
- Next.js
- shadcn/ui

### Backend

- Go
- ogen
- REST API
- OpenAPI
- gRPC
- dockertest
- sqlc
- MySQL
- Docker
- redis

### DevOps

- GitHub Actions
- ArgoCD

### Infrastructure

- Proxmox
- Ubuntu
- Terraform
- Ansible
- Kubernetes
- Prometheus
- Grafana
- Loki
- Fluent Bit
- cert-manager
- Let's Encrypt
