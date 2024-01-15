include .env

PROTO_FILE = proto
PROTO_OUT_DIR = src/app/grpc/proto

test:
	@echo "Running test command..."

run:
	go run main.go

gen-protoc:
	@echo "Delete generated protoc..."
	rm -rf $(PROTO_OUT_DIR)/*.go
	@echo "Generating protoc output..."
	protoc --proto_path=$(PROTO_FILE) --go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative \
    --go-grpc_out=$(PROTO_OUT_DIR) --go-grpc_opt=paths=source_relative \
    $(PROTO_FILE)/*.proto
	@echo "Finish Generating protoc..."