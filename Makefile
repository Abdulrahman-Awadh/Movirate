generate_grpc_code:
	protoc --go_out=movie --go_opt=paths=source_relative --go-grpc_out=movie --go-grpc_opt=paths=source_relative movie.proto
