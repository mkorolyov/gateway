run:
	GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info go run \
    	-ldflags "-X github.com/mkorolyov/core/discovery/consul.env=dev -X github.com/mkorolyov/core/discovery/consul.ip=127.0.0.1 -X github.com/mkorolyov/core/discovery/consul.port=9093" \
    	server/server.go