command to compile grpc
protoc -I ecommerce ecommerce/product_info.proto --go_out=plugins=grpc:ecommerce
