run:
	GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info go run \
    	-ldflags "-X github.com/mkorolyov/core/discovery/consul.env=dev -X github.com/mkorolyov/core/discovery/consul.ip=0.0.0.0 -X github.com/mkorolyov/core/discovery/consul.port=8080" \
    	server/server.go