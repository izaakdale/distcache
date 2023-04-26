CONFIG_PATH=${HOME}/.distcache

run:
	GRPC_PORT=5001 \
	REDIS_ADDR=localhost:6379 \
	BIND_ADDR=127.0.0.1 \
	BIND_PORT=7777 \
 	ADVERTISE_ADDR=127.0.0.1 \
	ADVERTISE_PORT=7777 \
	NAME=a1 \
	go run .

docker_run:
	docker run -d -p 5001:5001 \
	-e GRPC_PORT=5001 \
	-e REDIS_ADDR=redis:6379 \
	-e DB=0 \
	-e RECORD_TTL=0 \
	--network backend \
	--name distcache distcache

docker_run_serf1:
	docker run -d -p 6666:6666 -p 6666:6666/udp \
	-e BIND_ADDR=127.0.0.1 \
	-e BIND_PORT=6666 \
	-e ADVERTISE_ADDR=127.0.0.1 \
	-e ADVERTISE_PORT=6666 \
	-e CLUSTER_ADDR=127.0.0.1 \
	-e CLUSTER_PORT=6666 \
	-e NAME=a1 \
	--network host \
	--name serfer1 serfer
docker_run_serf2:
	docker run -d -p 6667:6667 -p 6667:6667/udp \
	-e BIND_ADDR=127.0.0.1 \
	-e BIND_PORT=6667 \
	-e ADVERTISE_ADDR=127.0.0.1 \
	-e ADVERTISE_PORT=6667 \
	-e CLUSTER_ADDR=127.0.0.1 \
	-e CLUSTER_PORT=6666 \
	-e NAME=a2 \
	--network host \
	--name serfer2 serfer

docker:
	docker build -t distcache .

.PHONY: gproto
gproto:
	protoc api/v1/*.proto \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--proto_path=.


init:
	mkdir -p ${CONFIG_PATH}

.PHONY: clean
clean:
	rm -rf ${CONFIG_PATH}
	mkdir -p ${CONFIG_PATH}

.PHONY: gencert
gencert:
	cfssl gencert -initca cert/ca-csr.json | cfssljson -bare ca
	cfssl gencert -ca=ca.pem \
	-ca-key=ca-key.pem \
	-config=cert/ca-config.json \
	-profile=server \
	cert/server-csr.json | cfssljson -bare server
	cfssl gencert -ca=ca.pem \
	-ca-key=ca-key.pem \
	-config=cert/ca-config.json \
	-profile=client \
	-cn="root" \
	cert/client-csr.json | cfssljson -bare client
	mv *.pem *csr ${CONFIG_PATH}

secret:
	kubectl create secret tls server-creds --cert=${CONFIG_PATH}/server.pem --key=${CONFIG_PATH}/server-key.pem
	kubectl create secret tls client-creds --cert=${CONFIG_PATH}/client.pem --key=${CONFIG_PATH}/client-key.pem
	kubectl create secret tls ca-creds --cert=${CONFIG_PATH}/ca.pem --key=${CONFIG_PATH}/ca-key.pem
