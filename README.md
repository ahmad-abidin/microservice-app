# Simple Microservice : e-commerce (not completed yet)
View Design [Here](./design/README.md)

generate proto angular and go:
protoc \
    --go_out=plugins=grpc:./[directory_on_server] \
    --plugin=protoc-gen-ts=./[directory_on_client(web)]/node_modules/.bin/protoc-gen-ts \
    --ts_out=service=true:./[directory_on_client(web)]/src/app \
    --js_out=import_style=commonjs,binary:./[directory_on_client(web)]/src/app \
    ./proto/[file_proto.proto]