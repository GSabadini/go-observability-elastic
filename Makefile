##for ((i=1;i<=100;i++)); do curl -v --header "Connection: keep-alive" "localhost:3000/ping"; done
##n ?= 50
##Example request n=100 p=3001 r=/query/test
request:
	n=$(n); \
	while [ $${n} -gt 0 ] ; do \
		curl -v --header "Connection: keep-alive" "http://localhost:$(p)$(r)"; \
		n=`expr $$n - 1`; \
	done; \
	true

up:
	docker-compose up -d elasticsearch kibana apm-server heartbeat postgres redis mysql

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