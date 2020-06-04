##for ((i=1;i<=100;i++)); do curl -v --header "Connection: keep-alive" "localhost:3000/ping"; done
n ?= 50
request:
	n=$(n); \
	while [ $${n} -gt 0 ] ; do \
		curl -v --header "Connection: keep-alive" "http://localhost:3001/query/gabriel"; \
		n=`expr $$n - 1`; \
	done; \
	true

up:
	docker-compose up -d elasticsearch kibana apm-server heartbeat

up-go-app-build:
	docker-compose up -d --build go-app

up-go-app:
	docker-compose up -d go-app

up-go-info-build:
	docker-compose up -d --build go-info

up-go-info:
	docker-compose up -d go-info

up-apps:
	docker-compose up -d go-info go-app

up-metric:
	docker-compose up -d metricbeat

up-heart:
	docker-compose up -d heartbeat

down:
	docker-compose down --remove-orphans