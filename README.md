# protoc-gen-govalidate
protoc plugin to generate polyglot message validators . it uses [go-playground](https://github.com/go-playground/validator) to generate validators for go.


## Installation
    
```bash 
go install github.com/aesoper101/protoc-gen-govalidate/cmd/protoc-gen-govalidate@latest
```

## Usage

```shell
protoc -I . -I path/to/validate/  --go_out=paths=source_relative:. --govalidate_out=paths=source_relative:. example/*.proto
```

## Dependencies

## Contributing