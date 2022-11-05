package main

import (
	"bytes"
	"embed"
	"github.com/aesoper101/go-utils/templatex"
	"google.golang.org/protobuf/compiler/protogen"
)

//go:embed template.tmpl
var templateFs embed.FS

type ruleInfo struct {
	Name       string
	FieldRules []fieldRule
	GoIdent    protogen.GoIdent
}

type fieldRule struct {
	Name          string
	Value         string
	GoPackageName string
}

type ruleInfoWrapper struct {
	Rules map[string]*ruleInfo
}

func (r *ruleInfoWrapper) execute() string {
	template, err := templatex.New("template.tmpl").ParseFS(templateFs, "template.tmpl")
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err := template.Execute(&buf, r); err != nil {
		panic(err)
	}

	return buf.String()
}
