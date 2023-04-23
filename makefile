CONFIG_PATH=${HOME}/.distcache

run:
	GRPC_PORT=5001 REDIS_ADDR=localhost:6379 go run .

docker_run:
	docker run -d -p 5001:5001 \
	-e GRPC_PORT=5001 \
	-e REDIS_ADDR=redis:6379 \
	-e DB=0 \
	-e RECORD_TTL=0 \
	--network backend \
	--name distcache distcache

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
	cfssl gencert -initca test/ca-csr.json | cfssljson -bare ca
	cfssl gencert -ca=ca.pem \
	-ca-key=ca-key.pem \
	-config=test/ca-config.json \
	-profile=server \
	test/server-csr.json | cfssljson -bare server
	cfssl gencert -ca=ca.pem \
	-ca-key=ca-key.pem \
	-config=test/ca-config.json \
	-profile=client \
	-cn="root" \
	test/client-csr.json | cfssljson -bare client
	mv *.pem *csr ${CONFIG_PATH}

secret:
	kubectl create secret tls distcache-security --cert=/home/izaakdale/.distcache/server.pem --key=/home/izaakdale/.distcache/server-key.pem
