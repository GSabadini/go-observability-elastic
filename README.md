# go-apm-elastic


- APM metrics (APM Server)
- Golang metrics (Metricbeat)
- Docker metrics (Metricbeat)
- System metrics (Metricbeat)
- Health check (Heartbeat)

Iniciar kibana, elasticsearch e apm-server
```sh
make up
```
Aguarde o Kibana iniciar...

Iniciar app
```sh
make up-app
```

Iniciar Heartbeat
```sh
make up-heart
```

Iniciar Metricbeat
```sh
make up-metric
```