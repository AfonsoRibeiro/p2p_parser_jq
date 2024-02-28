container_name="p2p_parser_jq"
image_name="p2p_parser_jq"
image_version="0.0.1"

main: build_go
	./${container_name} --log_level=debug --source_allow_insecure_connection=true --dest_allow_insecure_connection=true

curl:
	curl localhost:7700/metrics

build_go:
	go build -C src/main -o ../../p2p_parser_jq

pprof: build_go
	./${container_name} --pprof_on=true --log_level=debug

run_container: build_cache
	docker run -d --rm \
		--env SOURCE_PULSAR=pulsar://localhost:6650 \
		--env SOURCE_TOPIC=non-persistent://public/functions/clog \
		--env ALARMISTIC_DIR=/app/alarmistic/ \
		-v ./alarmistic/:/app/alarmistic/:z \
		-p 7700:7700 \
		--name ${container_name} \
		${image_name}:${image_version}

# workaround for dockerfile context
begin_build:
	mkdir -p build/gojq_extention build/p2p_parser
	cp -r ../p2p_parser/go.* build/gojq_extention/
	cp -r ../gojq_extention/src build/gojq_extention/src
	cp -r ../p2p_parser/go.* build/p2p_parser/
	cp -r ../p2p_parser/src build/p2p_parser/src

end_build:
	rm -r build/

build: begin_build
	echo "Building ${image_name}:${image_version} --no-cache"
	docker build -t ${image_name}:${image_version} . --no-cache
	make end_build

build_cache: begin_build
	echo "Building ${image_name}:${image_version} --with-cache"
	docker build -t ${image_name}:${image_version} .
	make end_build

docker_hub: build
	docker tag ${image_name}:${image_version} xcjsbsx/${image_name}:${image_version}
	docker push xcjsbsx/${image_name}:${image_version}

start_pulsar:
	docker run -d --rm --name pulsar --net elastic -p 6650:6650 -p 8080:8080 apachepulsar/pulsar:latest bin/pulsar standalone

launch_pprof:
	/home/afonso_sr/go/bin/pprof -http=:8081 myprogram.prof

test_jq:
	./private_filters/jq_scripts_dev/testjq.sh

clean:
	rm ${container_name}; \
	rm pprof/2023*; \
	rm -r build