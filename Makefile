clean:
	@echo "--- cleanup all build and generated files ---"
	@rm -vf app/infrastructure/proto/pb/*.pb.go

protoc: clean
	@echo "--- preparing proto output directories ---"
	@mkdir -p app/infrastructure/proto/pb

	@echo "--- Compiling all proto files ---"
	@cd ./app/infrastructure/proto && protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative *.proto

setup:
	@echo " --- Setup and generate configuration --- "
	@cp config/examples/server.yml.example config/server/server.yml
	@cp config/examples/database.yml.example config/database/database.yml

build: setup protoc
	@echo "--- Building binary file ---"
	@go build -o ./main cmd/server/oauth.go

build-rest: setup protoc
	@echo "--- Building binary file ---"
	@go build -o ./main cmd/rest/rest.go

build-grpc: setup protoc
	@echo "--- Building binary file ---"
	@go build -o ./main cmd/grpc/grpc.go

run-grpc:
	@go run cmd/grpc/grpc.go

migrate:
	@go run cmd/migrate/migrate.go

wire:
	@echo "-- generate dependency injection --"
	@wire ./app/infrastructure/container/