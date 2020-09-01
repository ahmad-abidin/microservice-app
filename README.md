# Simple Microservice : e-commerce (not completed yet)
View Design [Here](./design/README.md)

tools :
Angular :
[@improbable-eng/grpc-web](https://www.npmjs.com/package/@improbable-eng/grpc-web)
[ts-protoc-gen](https://www.npmjs.com/package/ts-protoc-gen)

generate proto angular and go:
protoc \
    --go_out=plugins=grpc:./auth-service \
    --plugin=protoc-gen-ts=./e-commerce-web/node_modules/.bin/protoc-gen-ts \
    --ts_out=service=true:./e-commerce-web/src/app \
    --js_out=import_style=commonjs,binary:./e-commerce-web/src/app \
    ./model/auth.proto