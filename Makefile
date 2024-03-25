build:
	docker build -t mwitamike/savannah:1.0.2 .

push:
	docker push mwitamike/savannah:1.0.2


mockserver:
	prism mock -h localhost -p 4010 docs/Ordermanagement.yaml


#Nomad

start-agent:
	sudo nomad agent -dev
status:
	nomad status




jaeger:
	docker run -d --name jaeger-instance \
    -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
    -p 5776:5775/udp \
    -p 6833:6831/udp \
    -p 6834:6832/udp \
    -p 5779:5778 \
    -p 16687:16686 \
    -p 14269:14268 \
    -p 9412:9411 \
    jaegertracing/all-in-one:latest




prom:
	docker pull prom/prometheus
run-prom:
	docker run --name prometheus -d -p 9090:9090 prom/prometheus

grafana:
	docker pull grafana/grafana
run-grafana:
	docker run --name grafana -d -p 3000:3000 grafana/grafana

