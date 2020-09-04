# Simple Microservice : e-commerce (not completed yet)
View Design [Here](./design/README.md)

tools :<br>
Angular :<br>
[@improbable-eng/grpc-web](https://www.npmjs.com/package/@improbable-eng/grpc-web)<br>
[ts-protoc-gen](https://www.npmjs.com/package/ts-protoc-gen)<br>
Golang :<br>
[mysql](https://github.com/go-sql-driver/mysql)<br>
[grpc](https://google.golang.org/grpc)<br>
[protoc-gen-go](https://github.com/golang/protobuf/protoc-gen-go)<br>
[go module]()<br>
other :<br>
[docker]()<br>
[docker-compose]()<br>

code architecture :
[clean architecture (uncle bob)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
[clean architecture (iman tumorang)](https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047)

generate proto angular and go:
protoc \
    --go_out=plugins=grpc:./auth-service/delivery/grpc \
    --plugin=protoc-gen-ts=./e-commerce-web/node_modules/.bin/protoc-gen-ts \
    --ts_out=service=true:./e-commerce-web/src/app \
    --js_out=import_style=commonjs,binary:./e-commerce-web/src/app \
    ./proto/auth.proto