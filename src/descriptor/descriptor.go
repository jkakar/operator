// Package generator normalizes user-defined protobuf files and normalizes RPC
// service descriptions into data structures suitable for code generation. This
// is used the protoc-gen-hubot, protoc-gen-cmd, and protoc-gen-operatord
// protobuf compiler plugins.
package descriptor

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/sr/operator/src/operator"
)

const undocumentedPlaceholder = "Undocumented."

type ServiceSet struct {
	Options  Options
	Services []*Service
}

type Options struct {
	BinaryName string
	Imports    []string
}

type Service struct {
	Name        string
	Description string
	Methods     []*Method
}

type Method struct {
	Name        string
	Description string
	Arguments   []*Argument
}

type Argument struct {
	Name        string
	Description string
}

func Describe(request *plugin.CodeGeneratorRequest) (*ServiceSet, error) {
	numFiles := len(g.request.FileToGenerate)
	if numFiles == 0 {
		return nil, errors.New("no file to generate")
	}
	params := make(map[string]string)
	for _, p := range strings.Split(g.request.GetParameter(), ",") {
		if i := strings.Index(p, "="); i < 0 {
			params[p] = ""
		} else {
			params[p[0:i]] = p[i+1:]
		}
	}
	binaryName := defaultBinaryName
	if val, ok := params["binary"]; ok {
		binaryName = val
	}
	numServices := 0
	for _, file := range g.request.ProtoFile {
		numServices = numServices + len(file.Service)
	}
	serviceSet := &ServiceSet{
		Options: &Options{
			BinaryName: binaryName,
			Imports:    make([]string, numFiles),
		},
		Services: make([]*Service, numServices),
	}
	for i, file := range g.request.FileToGenerate {
		// TODO(sr) substitute path.Ext() or whatever
		b := fmt.Sprintf("/%s", path.Base(file))
		serviceSet.Imports[i] = fmt.Sprintf("github.com/sr/operator/src/%s", strings.Replace(file, b, "", 1))
	}
	i := 0
	for _, file := range g.request.ProtoFile {
		messagesByName := make(map[string]*descriptor.DescriptorProto)
		for _, message := range file.MessageType {
			messagesByName[message.GetName()] = message
		}
		for _, service := range file.Service {
			if service.Options == nil {
				return nil, fmt.Errorf("options name for service %s is missing", service.GetName())
			}
			name, err := proto.GetExtension(service.Options, operator.E_Name)
			if err != nil {
				return nil, err
			}
			nameStr := *name.(*string)
			serviceSet.Services[i] = &Service{
				Name:        nameStr,
				BinaryName:  binaryName,
				Description: undocumentedPlaceholder,
				Methods:     make([]*methodDescriptor, len(service.Method)),
			}
			for j, method := range service.Method {
				input := messagesByName[strings.Split(method.GetInputType(), ".")[2]]
				serviceSet.Services[i].Methods[j] = &Method{
					Name:        method.GetName(),
					Description: undocumentedPlaceholder,
					Arguments:   make([]*Argument, len(input.Field)),
				}
				for k, field := range input.Field {
					serviceSet.Services[i].Methods[j].Arguments[k] = &Argument{
						// TODO(sr) deal with ID => Id etc better
						Name:        strings.Replace(field.GetName(), "ID", "Id", 1),
						Description: undocumentedPlaceholder,
					}
				}
			}
			i = i + 1
		}
		for _, loc := range file.GetSourceCodeInfo().GetLocation() {
			if loc.LeadingComments == nil {
				continue
			}
			if len(loc.Path) == 2 && loc.Path[0] == 6 {
				serviceSet.Services[loc.Path[1]].Description = loc.LeadingComments
			} else if len(loc.Path) == 4 && loc.Path[0] == 6 && loc.Path[2] == 2 {
				serviceSet.Services[loc.Path[1]].Methods[loc.Path[3]].Description = *loc.LeadingComments
			}
		}
	}
	return serviceSet, nil
}
