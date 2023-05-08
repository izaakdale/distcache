CONFIG_PATH=${HOME}/.distcache

run:
	GRPC_PORT=5001 \
	REDIS_ADDR=localhost:6379 \
	BIND_ADDR=127.0.0.1 \
	BIND_PORT=7777 \
 	ADVERTISE_ADDR=127.0.0.1 \
	ADVERTISE_PORT=7777 \
	NAME=a1 \
	CONFIG_DIR=/Users/izaakdale/.distcache \
	ENV=local \
	go run .

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
