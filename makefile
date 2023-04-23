run:
	GRPC_PORT=5001 GRPC_HOST=localhost REDIS_ADDR=localhost:6379 DB=0 RECORD_TTL=0 go run .

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
