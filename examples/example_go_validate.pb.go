// Code generated by protoc-gen-go-validate. DO NOT EDIT.

package example

import (
	context "context"
	validator "github.com/aesoper101/protoc-gen-govalidate/validator"
)

var (
	_exampleRule = map[string]string{
		"A": "required",
		"B": "required,gt=0",
	}
	_exampleNestedRule = map[string]string{
		"B": "required",
	}
)

func init() {
	validator.Validator().RegisterStructValidationMapRules(_exampleRule, Example{})
	validator.Validator().RegisterStructValidationMapRules(_exampleNestedRule, Example_Nested{})
}

func (x *Example) Validate() error {
	return validator.Validator().Struct(x)
}

func (x *Example) ValidateCtx(ctx context.Context) error {
	return validator.Validator().StructCtx(ctx, x)
}

func (x *Example_Nested) Validate() error {
	return validator.Validator().Struct(x)
}

func (x *Example_Nested) ValidateCtx(ctx context.Context) error {
	return validator.Validator().StructCtx(ctx, x)
}
