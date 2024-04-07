# QueryChat

## Sequence Diagram

```mermaid
sequenceDiagram
  actor User
  participant Frontend
  box User System
    participant Agent
    participant User DB
  end
  box QueryChat System
    participant Backend
    participant LLM
    participant DB
  end

  User ->> Frontend: Send prompt
  Note over Frontend, Agent: REST API (VPN)
  Frontend ->> Agent: Request (prompt)
  Note over Agent, Backend: gRPC
  Agent ->> Backend: Request (dbType, schema, prompt)
  Backend ->> DB: Check permission
  DB ->> Backend: Result
  alt Unauthorized
    Backend ->> Agent: 401 Error
    Agent ->> Frontend: 401 Error
    Frontend ->> User: Error Message
  end
  Backend ->> LLM: Send tuned prompt
  LLM ->> Backend: Output
  Backend ->> Agent: SQL query
  Agent ->> User DB: Execute SQL query
  User DB ->> Agent: Result
  Agent ->> Frontend: Visualizable data
  Frontend ->> User: Visualized data
```

## Technologies

### Frontend

- TypeScript
- Bun
- Next.js
- shadcn/ui
- Recharts

### Agent

- Go
- ogen
- REST API
- OpenAPI
- gRPC
- Docker

### Backend

- Go
- ogen
- gRPC
- dockertest
- sqlc
- Docker
- redis

### LLM

- Python
- gRPC
- OpenAI

### Database

- MySQL

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

## Development Flow

Trunk-based development. Every commit is a release candidate.

Treat commit messages as Pull Request equivalents. Commit message format is as follows:

```plaintext
<type>(<issue title & id>): <description>

<body>
```

- Commit titles should begin with a lowercase letter
- Should add issue id to link the commit to the issue

For example:

```plaintext
fix(a display bug #20045): fix a display bug in the user list

The user list is not displayed correctly when the user has a long name...
```
