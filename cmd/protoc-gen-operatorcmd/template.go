package main

import "github.com/sr/operator/generator"

var mainTemplate = generator.NewTemplate("main-gen.go",
	`// Code generated by protoc-gen-operatorcmd
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/sr/operator"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	{{- range .Services}}
	{{.PackageName}} "{{.ImportPath}}"
	{{- end}}
)

const programName = "{{.Options.BinaryName}}"

var cmd = operator.NewCommand(
	programName,
	[]operator.ServiceCommand{
{{- range .Services}}
		{
{{- $serviceName := .Name }}
{{- $serviceFullName := .FullName }}
			Name:     "{{ $serviceName }}",
			Synopsis: `+"`"+`{{ .Description }}`+"`"+`,
			Methods: []operator.MethodCommand{
	{{- range .Methods}}
				{
					Name: "{{dasherize .Name}}",
					Synopsis: `+"`"+`{{.Description}}`+"`"+`,
					Run: func(ctx *operator.CommandContext) (string, error) {
			{{- range .Arguments}}
						{{.Name}} := ctx.Flags.String("{{dasherize .Name}}", "", "")
			{{- end}}
						if err := ctx.Flags.Parse(ctx.Args); err != nil {
							return "", err
						}
						conn, err := dial(ctx.Address)
						if err != nil {
							return "", err
						}
						defer func() { _ = conn.Close() }()
						client := {{$serviceName}}.New{{$serviceFullName}}Client(conn)
						response, err := client.{{.Name}}(
							context.Background(),
							&{{$serviceName}}.{{.Input}}{
								Source: ctx.Source,
								{{- range .Arguments}}
								{{camelCase .Name}}: *{{.Name}},
								{{- end}}
							},
						)
						if err != nil {
							return "", err
						}
						return response.Output.PlainText, nil
					},
				},
	{{- end }}
			},
		},
{{end}}
	},
)

func main() {
	status, output := cmd.Run(os.Args)
	if status != 0 {
		if _, err := fmt.Fprintf(os.Stderr, "%s: %s\n", programName, output); err != nil {
			panic(err)
		}
	} else {
		if _, err := io.WriteString(os.Stdout, output); err != nil {
			panic(err)
		}
	}
	os.Exit(status)
}

func dial(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}`)
