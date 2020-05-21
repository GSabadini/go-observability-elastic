send-request:
	for ((i=1;i<=100;i++)); do curl -v --header "Connection: keep-alive" "localhost:3000/ping"; done

up:
	docker-compose up -d elasticsearch kibana apm-server app

up-metric:
	docker-compose up -d metricbeat