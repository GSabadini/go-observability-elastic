##for ((i=1;i<=100;i++)); do curl -v --header "Connection: keep-alive" "localhost:3000/ping"; done
n ?= 25
request:
	n=$(n); \
	while [ $${n} -gt 0 ] ; do \
		curl -v --header "Connection: keep-alive" "localhost:3000/ping"; \
		n=`expr $$n - 1`; \
	done; \
	true

up:
	docker-compose up -d elasticsearch kibana apm-server

up-app-build:
	docker-compose up -d --build app

up-app:
	docker-compose up -d app

up-metric:
	docker-compose up -d metricbeat app

up-heart:
	docker-compose up -d heartbeat app

down:
	docker-compose down --remove-orphans