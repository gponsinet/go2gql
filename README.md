go2gql
==============

Tool, which generates GraphQL Schema for [graphql-go](https://github.com/saturn4er/graphql) package based on external data sources

## Installation
```bash
$ go get github.com/EGT-Ukraine/go2gql/cmd/go2gql
```

## Usage
```
$ ./go2gql -c "<config path>"
```

## Generation process
### Plugins
The generation process is built around plugins. Plugin is a go type that implements interface

```go
type Plugin interface {
	Init(*GenerateConfig, []Plugin) error
	Prepare() error
	Name() string
	PrintInfo(info Infos)
	Infos() map[string]string
	Generate() error
}
```

1) reading config
2) plugins initialization ( calling Init() method of each plugin )
3) plugins preparation ( calling Prepare() method of each plugin )
3) plugins generation ( calling Generate() method of each plugin )

## Default plugins
Default plugins places in ./generator/plugins

### `graphql` plugin
`graphql` generates:
 - GraphQL Schema based on config and relying on data sent to him from other packages
 - GraphQL types and resolvers relying on data sent to him from other packages

#### config example
Here's a simple config example of how to generate such GraphQL schema
```GraphQL
type Query {
    fieldA FieldA
}
type FieldA {
    srvField SomeServiceObj
}
type SomeServiceObj {
    ... SomeService methods
}
type Mutation {
    ... SomeServiceMutations methods
}
```


```yml
...
...
graphql_schemas:
  - name: "ExampleSchema"                           # Schema name
    output_path: "./services_api/schema/auth.go"    # Where to put schema code
    output_package: "schema"                        # Package name
    queries:
      type: "OBJECT"
      fields:
        - field:       "fieldA"
          type:        "OBJECT"                     # Field type (OBJECT|SERVICE)
          object_name: "FieldA"
          fields:
            - field:        "srvField"
              object_name:  "SomeServiceObj"
              service:      "SomeService"           # Service name
              type:         "SERVICE"
              exclude_methods:
                - "excluded_method_1"
              filter_methods:
                - "excluded_method_1"
                - "excluded_method_3"

    mutations:
      # similary to queries
      type: "SERVICE"
      service: "SomeServiceMutations"
...
```


### `proto2gql` plugin
`proto2gql` plugin parses .proto files, defined in config and pass them to `graphql` plugin.

#### Service names
For each service of each proto file, `proto2gql` adds two services to `graphql` plugins with names:
 - `ServiceName` or `ServiceAlias`
 - `ServiceName`Mutations or `ServiceAlias`Mutations

#### Config example

```yml
...
...
proto2gql:
    output_path: "./proto_out/"      # path, where to put generated code
    paths:                           # path, where parser will search for imports
        - "./vendor"
        - "$GOPATH/src"
    import_aliases:                  # imports aliases for all files
        - google/protobuf/descriptor.proto: "github.com/golang/protobuf/protoc-gen-go/descriptor/descriptor.proto"
        - google/protobuf/timestamp.proto:  "github.com/golang/protobuf/ptypes/timestamp/timestamp.proto"
        - google/protobuf/empty.proto:      "github.com/golang/protobuf/ptypes/empty/empty.proto"
    files:
      # If you want to add some settings to imported .proto file, just add it here
      - name: "Name"                  # name of .proto file
        proto_path: "./a.proto"       # path to .proto file
        output_path: "./schema/name"  # file-specific path, where to put generated code
        proto_go_package: "github.com/gogo/protobuf/types" # Go package of .proto file grpc client (which was previously generated by `protoc`)
        gql_messages_prefix: "Msg"    # prefix, which will be added to all generated GraphQL Messages(including maps)
        gql_enums_prefix:    "Enm"    # prefix, which will be added to all generated GraphQL Enums
        paths:                        # file specific path, where parser will search for imports
          - "./a_proto_deps/"
        imports_aiases:               # file specific imports aliases
          - google/protobuf/timestamp.proto:  "github.com/gogo/protobuf/protobuf/google/protobuf/timestamp.proto"
        services:                     # services settings
          "serviceName":              # service name
            alias: "someAlias"        # service name alias
            methods:
              "methodName":           # method name
                alias: "methodAlias"
                request_type: "QUERY" # method type in GraphQL Schema (QUERY|MUTATION)
        messages:                     # messages settings
          - "Request$":               # message name match regex
              fields:
                "ip_address":
                  context_key: "ip"   # context key, where unmarshaller will get this field from
          - "Response$":
              error_field: "Error"    # name of payload error field
      - ...
      - ...


...
...
```

### `swagger2gql` plugin
`proto2gql` plugin parses swagger files, defined in config and pass them to `graphql` plugin.

#### Service names
For each tag of each swagger file, `proto2gql` adds two services to `graphql` plugins with names:
 - `pascalized(tag-name)` or `ServiceAlias`
 - `pascalized(tag-name)`Mutations or `ServiceAlias`Mutations

pascalized() here means, that it converts string to CamelCase format and removes all characters that is not valid in GraphQL Object name

#### Config example
```yml
...
...
swagger2gql:
    output_path: "./swagger_out/"           # Path, where to put generated code
    objects:                                # GraphQL objects configs
      - "Response$":                        # Object name
          error_field: "error"              # name of payload error field
      - "InputParams$":
          fields:                           # Object fields config
            ip:                             # field name
              context_key: "user_ip"        # context key, where unmarshaller will get this field from
    files:
      - name: "Swagger file number 1"         # swagger file name
        path: "./service1/swagger.json"       # path to swagger file
        models_go_path: "github.com/myproject/service1/client/models" # path to generated using goswagger models
        output_pkg: "gql_service1"            # output go package name
        output_path: "./service1/schema/"     # file-specific path, where to put generated code
        gql_objects_prefix: "Srv1"            # prefix, which will be added to all generated GraphQL Objects
        tags:                                 # tags settings
          "some-swagger-tag":                 # tag name
            client_go_package: "github.com/myproject/service1/client/some_swagger_tag/client"   # go client package
            service_name: "SomeSwaggerTag"    # tag service name
            methods:                          # tag method settings
              "/somes":                       # request path
                get:                          # request method (get/post/put/options...)
                  alias: "get"
                  request_type: "QUERY"       # request type (QUERY|MUTATIONS)
        objects:                              # file specific objects settings
         - "Request$":
           fields:
             "user_id":
               context_key: "user_id"
...
...
```


### Simple plugin example
Here's a simple plugin, that gets access to graphql plugin. Fetches information about services and their methods, and then, render them to .yml file

```go
package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/EGT-Ukraine/go2gql/generator"
	"github.com/EGT-Ukraine/go2gql/generator/plugins/graphql"
)

const Name = "gql_services_rules"

type plugin struct {
	gqlPlugin *graphql.Plugin
}

func (p *plugin) Init(_ *generator.GenerateConfig, plugins []generator.Plugin) error {
	for _, plugin := range plugins {
		if g, ok := plugin.(*graphql.Plugin); ok {
			p.gqlPlugin = g
			return nil
		}
	}
	return errors.New("graphql plugin was not found")
}

func (p plugin) Generate() error {
	file, err := os.OpenFile("./services_access.yml", os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return errors.Wrap(err, "failed to open services_access.yml")
	}
	defer file.Close()
	for _, typesFile := range p.gqlPlugin.Types() {
		for _, service := range typesFile.Services {
			if len(service.Methods) == 0 {
				continue
			}
			file.WriteString(service.Name + ":\n")
			for _, method := range service.Methods {
				file.WriteString("   " + method.Name + ":\n")
			}
		}
	}
	return nil
}
func (plugin) Name() string                   { return Name }
func (plugin) Prepare() error                 { return nil }
func (plugin) PrintInfo(info generator.Infos) {}
func (plugin) Infos() map[string]string       { return nil }

func Plugin() generator.Plugin {
	return new(plugin)
}
func main(){}
```