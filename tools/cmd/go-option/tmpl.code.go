package main

import (
	"strings"

	strings_ "github.com/searKing/golang/go/strings"
)

type TmplConfigRender struct {
	// Print the header and package clause.
	GoOptionToolName       string
	GoOptionToolArgs       []string
	GoOptionToolArgsJoined string

	PackageName string
	ImportPaths []string
	ValDecls    []string

	TargetTypeName               string // type name of target type
	TargetTypeImport             string // import path of target type
	TargetTypeGenericDeclaration string // the Generic type of the struct type
	TargetTypeGenericParams      string // the Generic params of the struct type
	TrimmedTypeName              string // trimmed type name of target type

	OptionInterfaceName string // option interface name of target type
	OptionStructName    string // option struct name of target type

	ApplyOptionsAsMemberFunction bool // ApplyOptions can be registered as OptionType's member function
}

func (t *TmplConfigRender) Complete() {
	t.GoOptionToolArgsJoined = strings.Join(t.GoOptionToolArgs, " ")
	t.ApplyOptionsAsMemberFunction = strings.TrimSpace(t.TargetTypeImport) == ""

	t.OptionInterfaceName = strings_.UpperCamelCaseSlice("option")
	t.OptionStructName = strings_.UpperCamelCaseSlice("config")
	if !*trim && t.TrimmedTypeName != "" {
		t.OptionInterfaceName = strings_.UpperCamelCaseSlice(t.TrimmedTypeName, "option")
		t.OptionStructName = strings_.UpperCamelCaseSlice(t.TrimmedTypeName, "config")
	}

	importPath := strings.TrimSpace(t.TargetTypeImport)
	if importPath != "" {
		t.ImportPaths = append(t.ImportPaths, importPath)
	}

	_, defaultValDecl := createValAndNameDecl(t.TargetTypeName)
	if defaultValDecl != "" {
		t.ValDecls = append(t.ValDecls, defaultValDecl)
	}
}

const tmplConfig = `// Code generated by "{{.GoOptionToolName}} {{.GoOptionToolArgsJoined}}"; EDIT IT ANYWAY.
// Install {{.GoOptionToolName}} by "go get install github.com/searKing/golang/tools/{{.GoOptionToolName}}"

package {{.PackageName}} 

import "fmt"

{{range $path := .ImportPaths}} import "{{$path}}" {{end}}

// {{.OptionStructName}} is config param as factory config
// Code borrowed from https://github.com/kubernetes/kubernetes
// call chains: New{{.OptionStructName}}{{.TargetTypeGenericDeclaration}} -> Complete{{.TargetTypeGenericParams}} -> Validate{{.TargetTypeGenericParams}} -> New{{.TargetTypeGenericParams}}|Apply{{.TargetTypeGenericParams}}
type {{.OptionStructName}}{{.TargetTypeGenericDeclaration}} struct {
	// TODO Add config fields here
}

{{- if .ApplyOptionsAsMemberFunction }}
type completed{{.OptionStructName}}{{.TargetTypeGenericDeclaration}} struct {
	*{{.OptionStructName}}{{.TargetTypeGenericParams}}

	//===========================================================================
	// values below here are filled in during completion
	//===========================================================================
}

// Completed{{.OptionStructName}} is config ready to use.
type Completed{{.OptionStructName}}{{.TargetTypeGenericDeclaration}} struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completed{{.OptionStructName}}{{.TargetTypeGenericParams}}
}

// New{{.OptionStructName}} returns a Config struct with the default values
func New{{.OptionStructName}}{{.TargetTypeGenericDeclaration}}() *{{.OptionStructName}}{{.TargetTypeGenericParams}} {
	// TODO Add default configs here
	return &{{.OptionStructName}}{{.TargetTypeGenericParams}}{}
}

// Complete fills in any fields not set that are required to have valid data and can be derived
// from other fields. If you're going to ApplyOptions, do that first. It's mutating the receiver.
func (o *{{.OptionStructName}}{{.TargetTypeGenericParams}}) Complete() Completed{{.OptionStructName}}{{.TargetTypeGenericParams}} {
	// TODO Add custom codes here
	return Completed{{.OptionStructName}}{{.TargetTypeGenericParams}}{&completed{{.OptionStructName}}{{.TargetTypeGenericParams}}{o}}
}

// Validate checks {{.OptionStructName}}{{.TargetTypeGenericParams}} and return a slice of found errs.
func (o *{{.OptionStructName}}{{.TargetTypeGenericParams}}) Validate() []error {
	var errs []error
	// TODO Add custom validate codes here
	return errs
}

// New creates a new server which logically combines the handling chain with the passed server.
// The handler chain in particular can be difficult as it starts delegating.
// New usually called after Complete
func (c completed{{.OptionStructName}}{{.TargetTypeGenericParams}}) New() (*{{.TargetTypeName}}{{.TargetTypeGenericParams}}, error) {
	// TODO Add custom codes here
	return nil, fmt.Errorf("not implemented")
}

// Apply set options and something else as global init, act likes New but without {{.OptionStructName}}'s instance
// Apply usually called after Complete
func (c completed{{.OptionStructName}}{{.TargetTypeGenericParams}}) Apply() error {
	// TODO Add custom codes here
	return fmt.Errorf("not implemented")
}

{{- else}}
// Complete fills in any fields not set that are required to have valid data and can be derived
// from other fields. If you're going to ApplyOptions, do that first. It's mutating the receiver.
func Complete{{.TargetTypeGenericDeclaration}}(o *{{.OptionStructName}}, options ...{{.OptionInterfaceName}}{{.TargetTypeGenericParams}}) Completed{{.OptionStructName}} {
	ApplyOptions(o, options...)
	// TODO Add custom codes here
	return Completed{{.OptionStructName}}{{.TargetTypeGenericParams}}{&completed{{.OptionStructName}}{{.TargetTypeGenericParams}}{o}}
}
{{- end}}
`
