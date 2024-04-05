# QueryChat

## Sequence Diagram

```mermaid
sequenceDiagram
  box User System
    actor User
    participant Frontend
    participant Agent
    participant User DB
  end
  box querychat
    participant Backend
    participant LLM
    participant DB
  end

  User ->> Frontend: Input prompt
  Note over Frontend, Agent: REST API
  Frontend ->> Agent: Request
  Note over Agent, Backend: gRPC
  Agent ->> Backend: Request
  Backend ->> DB: Check permission
  DB ->> Backend: Return result
  alt Unauthorized
    Backend ->> Agent: 401 Error
    Agent ->> Frontend: 401 Error
    Frontend ->> User: 401 Error
  end
  Backend ->> LLM: Send prompt
  LLM ->> Backend: Return SQL query
  Backend ->> Agent: Return SQL query
  Agent ->> User DB: Execute SQL query
  User DB ->> Agent: Return result
  Agent ->> Frontend: Return visualizable data
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

## Development Flow

Trunk-based development. Every commit is a release candidate.

Treat commit messages as Pull Request equivalents. Commit message format is as follows:

```plaintext
<type>(<issue title & id>): <description>

[body]
```

- Commit titles should begin with a lowercase letter
- Should add issue id to link the commit to the issue

For example:

```plaintext
fix(a display bug #20045): fix a display bug in the user list

The user list is not displayed correctly when the user has a long name...
```
