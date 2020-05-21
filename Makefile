##for ((i=1;i<=100;i++)); do curl -v --header "Connection: keep-alive" "localhost:3000/ping"; done
n ?= 10
request:
	n=$(n); \
	while [ $${n} -gt 0 ] ; do \
		curl -v --header "Connection: keep-alive" "localhost:3000/ping"; \
		n=`expr $$n - 1`; \
	done; \
	true

up:
	docker-compose up -d --build elasticsearch kibana apm-server app

up-metric:
	docker-compose up -d metricbeat