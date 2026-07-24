# RAG Security Proxy

An open-source access-control proxy for Retrieval-Augmented Generation (RAG) pipelines.

It sits between your application/LLM orchestrator and a vector database, and enforces
per-user authorization on retrieval results: **unauthorized chunks never reach the LLM context**.

## Why

RAG pipelines typically query a vector DB with a service account, so every user can
retrieve every indexed document. This proxy evaluates access policies (subject → resource)
on each retrieval and filters out chunks the caller is not allowed to see.

## Status

Early skeleton. Only the project structure and a health endpoint exist.
No policy evaluation or retrieval proxying is implemented yet.

## Architecture (target)

```
client ──(JWT from Keycloak)──▶ rag-security-proxy ──▶ vector DB (Milvus)
                                      │
                                      ├── policy evaluation (per chunk / per collection)
                                      └── audit log (Postgres)
```

## Project layout

```
cmd/server/       entrypoint (HTTP server)
internal/health/  liveness/readiness handlers
internal/policy/  policy evaluation (interfaces only for now)
Dockerfile        multi-stage build, distroless runtime
docker-compose.yml local dev stack: Keycloak, Milvus, Postgres, the proxy
Makefile          build / test / run shortcuts
```

## Quick start

Requirements: Go 1.24+, Docker with Compose.

```bash
# run the server locally
make run
curl http://localhost:8080/healthz
# {"status":"ok"}

# or the full local stack (Keycloak, Milvus, Postgres, proxy)
make compose-up
```

Local endpoints:

| Service  | URL                     | Credentials        |
|----------|-------------------------|--------------------|
| Proxy    | http://localhost:8080   | —                  |
| Keycloak | http://localhost:8081   | admin / admin      |
| Milvus   | localhost:19530 (gRPC)  | —                  |
| Postgres | localhost:5432          | rag / rag          |

## Roadmap

1. JWT validation middleware (Keycloak OIDC).
2. Policy model + `internal/policy` implementation (Postgres-backed).
3. Retrieval proxy endpoint in front of Milvus with per-chunk filtering.
4. Audit logging.
5. Integration tests against the compose stack.

## License

Apache-2.0 (planned).
