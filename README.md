# protoc-gen-govalidate
protoc plugin to generate polyglot message validators . it uses [go-playground](https://github.com/go-playground/validator) to generate validators for go.


## Installation
go install github.com/aesoper101/protoc-gen-govalidate/cmd/protoc-gen-govalidate@latest

## Usage
protoc --proto_path=path/to/your/proto --go_out=paths=source_relative:. --govalidate_out=paths=source_relative:. *.proto

## Dependencies

## Contributing