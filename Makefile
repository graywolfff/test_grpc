gen:
	protoc --go_out=coffeeshop_protos --go_opt=paths=source_relative --go-grpc_out=coffeeshop_protos --go-grpc_opt=paths=source_relative coffee_shop.proto
clean:
	rm -rf coffeeshop_proto/*