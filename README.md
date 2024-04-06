# QueryChat

## Sequence Diagram

```mermaid
sequenceDiagram
  actor User
  box User System
    participant Frontend
    participant Agent
    participant User DB
  end
  box QueryChat System
    participant Backend
    participant LLM
    participant DB
  end

  User ->> Frontend: Input prompt
  Note over Frontend, Agent: REST API
  Frontend ->> Agent: Request (prompt)
  Note over Agent, Backend: gRPC
  Agent ->> Backend: Request (prompt, schema)
  Backend ->> DB: Check permission
  DB ->> Backend: Return result
  alt Unauthorized
    Backend ->> Agent: 401 Error
    Agent ->> Frontend: 401 Error
    Frontend ->> User: Error Message
  end
  Backend ->> LLM: Send tuned prompt
  LLM ->> Backend: Return SQL query
  Backend ->> Agent: Return SQL query
  Agent ->> User DB: Execute SQL query
  User DB ->> Agent: Return result
  Agent ->> Frontend: Return visualizable data
  Frontend ->> User: Visualized data
```

## Technologies

### Frontend

- TypeScript
- Next.js
- shadcn/ui

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
- FastAPI
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
