<h1 align="center">🚦 DThrottlr : A Distributed Rate Limiter</h1>
<p align="center"><i>Built in Go - To learn Go.</i></p>

<p align="center">
<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white">
<img src="https://img.shields.io/badge/Lua-2C2D72?style=for-the-badge&logo=lua&logoColor=white">
<img src="https://img.shields.io/badge/chi_router-3E3E3E?style=for-the-badge">
<img src="https://img.shields.io/badge/go--redis/v9-3E3E3E?style=for-the-badge">
<img src="https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white">
<img src="https://img.shields.io/badge/Chart.js-FF6384?style=for-the-badge&logo=chartdotjs&logoColor=white">
<img src="https://img.shields.io/badge/WebSockets-black?style=for-the-badge&logo=socketdotio&logoColor=white">
<img src="https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white">
<img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white">
<img src="https://img.shields.io/badge/NGINX-009639?style=for-the-badge&logo=nginx&logoColor=white">
<img src="https://img.shields.io/badge/Kubernetes-326CE5?style=for-the-badge&logo=kubernetes&logoColor=white">
</p>

<p align="justify">
A distributed rate limiter built in Go, designed to run across multiple stateless application nodes sharing a single Redis backend. This project uses the token bucket algorithm, with all rate limit checks executed atomically inside Redis using Lua scripting, so there is no need for distributed locks.

I'm building this project mainly to get hands on with Go, since I come from a background of about three years of backend development but haven't worked with Go in production before. The plan is to build it phase by phase, learning the language and the distributed systems concepts along the way.
</p>

## Why this project

<p align="justify">
Rate limiting sounds simple until you have more than one server. Once you scale horizontally, each node needs to agree on how many requests a client has made, and that means the state has to live somewhere shared. Redis with Lua scripting solves this cleanly: the check-and-decrement operation happens in a single atomic step on the Redis server itself, so race conditions between nodes are avoided by design rather than by locking.
</p>

## Features (planned)

<ul style="list-style-type:square">
    <li>Distributed architecture with shared global rate limit state across N nodes</li>
    <li>Token bucket algorithm supporting burst traffic up to a configured capacity</li>
    <li>Atomic Lua scripts for race condition free evaluation, no distributed locks</li>
    <li>Single YAML config file to tune capacity, refill rate, client count and ports</li>
    <li>Real time dashboard showing allowed vs throttled traffic</li>
    <li>Docker Compose setup with a path to Kubernetes Deployments and HPAs</li>
</ul>

## Progress

<p align="justify">
This project is being built in phases. Here is where things stand right now.
<ul style="list-style-type:square">
    <li> [x] Phase 1: Server skeleton — Go module initialized, chi router set up, health check endpoint working.</li>
    <li> [x] Phase 2: Config loader — `config.yaml` drives server port, Redis connection, bucket capacity, refill rate and node settings, parsed and validated on startup.</li>
    <li> [x] Phase 3: Redis + Lua token bucket — Atomic token bucket logic written in Lua, embedded into the Go binary, executed via go-redis. Verified with unit tests and manual runs, including refill behavior and TTL based cleanup of idle client buckets.</li>
    <li> [x] Phase 4: HTTP middleware integration — Rate limiting wired in as chi middleware, returns 429 with a Retry-After header when a client is throttled, tested end to end against a live endpoint.</li>
    <li> [ ] Phase 5: Multi node simulation — Run multiple instances of the server against the same Redis and confirm state is shared correctly.</li>
    <li> [ ] Phase 6: Metrics and dashboard backend — Track allowed and throttled counts, expose over WebSocket or SSE.</li>
    <li> [ ] Phase 7: Dashboard frontend — Chart.js based live view of traffic with simulation controls.</li>
    <li> [ ] Phase 8: Load testing — Configurable client simulator to generate load and validate behavior under pressure.</li>
    <li> [ ] Phase 9: Dockerize — Multi stage builds, Compose setup with app nodes, Redis, NGINX and dashboard.</li>
    <li> [ ] Phase 10: Kubernetes — Deployment, Service and HPA manifests.</li>
    <li> [ ] Phase 11: Polish — Tests, linting, structured logging, graceful shutdown, documentation.</li>
</ul>
</p>

## Project Structure

```
dthrottlr/
├── cmd/
│   └── server/          # entrypoint
├── internal/
│   ├── config/          # config.yaml loader
│   ├── limiter/         # token bucket logic and embedded Lua script
│   ├── redisclient/     # redis connection setup
│   ├── middleware/      # rate limit middleware
│   └── server/          # router and handlers
├── dashboard/           # static frontend (planned)
├── deploy/
│   ├── docker/
│   ├── nginx/
│   └── k8s/
├── config.yaml
└── go.mod
```

## Running locally

Start Redis:

```bash
docker run -d -p 6379:6379 redis
```

Run the server:

```bash
go run cmd/server/main.go
```

Hit the rate limited endpoint:

```bash
curl -i localhost:<port>/ping
```

<p align="justify">
You should see 200 responses until the bucket is depleted, then 429 Too Many Requests until it refills.
</p>

## License

MIT License
