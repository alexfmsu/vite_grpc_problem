protoc 												\
	-I proto proto/*.proto 							\
	--js_out=import_style=commonjs:./proto \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./proto
