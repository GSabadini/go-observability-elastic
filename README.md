# go-apm-elastic

- APM metrics (APM Server)
- Golang metrics (Metricbeat)
- Docker metrics (Metricbeat)
- System metrics (Metricbeat)
- Redis metrics (Metricbeat)
- Health check (Heartbeat)

## go-info

- Start app in port 3000
- depends_on: postgres

```sh
make up-go-info-build
```
```sh
make up-go-info
```

| Endpoint        | HTTP Method             | Description            |
| --------------- | :---------------------: | :-----------------:    |
| `/info`         | `GET`                   | `Info network`         |
| `/health`       | `GET`                   | `Healthcheck`          |
| `/query/{name}`        | `GET`                   | `Query in Postgres` |
| `/http-external`| `GET`                   | `External HTTP request for go-app` |
| `/debug/vars`   | `GET`                   | `Metrics  for golang metrics` |

## go-app

- Start app in port 3001
- depends_on: redis
```sh
make up-go-info-build
```
```sh
make up-go-info
```

| Endpoint        | HTTP Method             | Description            |
| --------------- | :---------------------: | :-----------------:    |
| `/ping`         | `GET`                   | `Info network`         |
| `/health`       | `GET`                   | `Healthcheck`          |
| `/query/{name}`        | `GET`                   | `Query in SQLite` |
| `/cache/{key}`        | `GET`                   | `Find in Redis` |
| `/http-external`| `GET`                   | `External HTTP request for go-info` |
| `/debug/vars`   | `GET`                   | `Metrics for golang metrics` |


---

- Start Kibana, Elasticsearch, APM-Server and Heartbeat
```sh
make up
```

- Start Metricbeat after starting others containers
```sh
make up-metric
```

- Send request
```sh
make request PORT PATH
```

- Kill all containers
```sh
make down
```
