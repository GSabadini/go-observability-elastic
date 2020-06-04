# Go Observability Elastic

1. **APM Server**
    - Web requests
    - SQL queries
    - Outgoing HTTP requests
    - Panic tracking
    - Custom spans

- **Metricbeat**
    - Golang Metrics
    - Docker Metrics
    - System Metrics
    - Redis Metrics

- **Heartbeat**
    - Elasticsearch
    - Kibana
    - Golang apss

---

### Step by step

- Start Kibana, Elasticsearch, APM-Server, Postgres, Redis and Heartbeat
```sh
make up
```
#### go-info

- Start app
```sh
make up-go-info
```
- listen port 3001
- depends_on: postgres

| Endpoint        | HTTP Method             | Description            |
| --------------- | :---------------------: | :-----------------:    |
| `/info`         | `GET`                   | `Info network`         |
| `/health`       | `GET`                   | `Healthcheck`          |
| `/query/{name}` | `GET`                   | `Query in Postgres` |
| `/http-external`| `GET`                   | `HTTP request for app go-app` |
| `/debug/vars`   | `GET`                   | `Metrics  for golang metrics` |

#### go-app

- Start app in port 3000
```sh
make up-go-info
```

- listen port 3000
- depends_on: redis


| Endpoint        | HTTP Method             | Description            |
| --------------- | :---------------------: | :-----------------:    |
| `/ping`         | `GET`                   | `Pong response`        |
| `/health`       | `GET`                   | `Healthcheck`          |
| `/query/{name}` | `GET`                   | `Query in SQLite`      |
| `/cache/{key}`  | `GET`                   | `Find in Redis`        |
| `/http-external`| `GET`                   | `HTTP request for app go-info` |
| `/debug/vars`   | `GET`                   | `Metrics for golang metrics` |


- Start Metricbeat after starting others containers
```sh
make up-metric
```

#### Others commands

- Build app go-info
```sh
make up-go-info-build
```

Build app go-app
```sh
make up-go-app-build
```

- Send request
```sh
make request PORT PATH
```

- Kill all containers
```sh
make down
```
