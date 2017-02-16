package main

import "github.com/sr/operator/generator"

var template = generator.NewTemplate("builder-gen.go",
	`// Code generated by protoc-gen-operatord
package main

import (
	"fmt"

	"github.com/sr/operator"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

{{range $k, $v := .Imports}}
	{{$k}} "{{$v}}"
{{end}}
)

func invoker(ctx context.Context, conn *grpc.ClientConn, req *operator.Request, pkg string) error {
{{- range .Services}}
	{{- $pkg := .Package }}
	{{- $svc := .Name }}
	if req.Call.Service == fmt.Sprintf("%s.{{.Name}}", pkg) {
	{{- range .Methods }}
		if req.Call.Method == "{{.Name}}" {
			client := {{$pkg}}.New{{$svc}}Client(conn)
			_, err := client.{{.Name}}(
				ctx,
				&{{$pkg}}.{{.Input}}{
					Request: req,
					{{- range .Arguments}}
					{{inputField .Name}}: req.Call.Args["{{.Name}}"],
					{{- end}}
				},
			)
			return err
		}
	{{- end }}
	}
{{- end }}
	return fmt.Errorf("no such service: `+"`"+"%s %s"+"`"+`", req.Call.Service, req.Call.Method)
}
`)