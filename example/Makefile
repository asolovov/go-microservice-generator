generate-proto:
	cd protocols/user && protoc --descriptor_set_out=descriptor.pb --go_out=paths=source_relative:. --go_opt=paths=source_relative  --go-grpc_out=paths=source_relative:. --go-grpc_opt=paths=source_relative  *.proto

tidy:
	go mod tidy