// This file was generated by github.com/EGT-Ukraine/go2gql. DO NOT EDIT IT
package schema

import (
	graphql "github.com/saturn4er/graphql"

	interceptors "github.com/EGT-Ukraine/go2gql/api/interceptors"
	example "github.com/EGT-Ukraine/go2gql/example/out/example"
	proto "github.com/EGT-Ukraine/go2gql/example/proto"
)

type ExampleSchemaSchemaClients struct {
	ServiceExampleClient          proto.ServiceExampleClient
	ServiceExampleMutationsClient proto.ServiceExampleClient
}

func GetExampleSchemaSchema(cls ExampleSchemaSchemaClients, ih *interceptors.InterceptorHandler) (graphql.Schema, error) {
	var ServiceExampleFields = example.GetServiceExampleServiceMethods(cls.ServiceExampleClient, ih)
	var _ = ServiceExampleFields
	var ServiceExampleMutationsFields = example.GetServiceExampleMutationsServiceMethods(cls.ServiceExampleMutationsClient, ih)
	var _ = ServiceExampleMutationsFields
	var Query = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getQueryMethod": ServiceExampleFields["getQueryMethod"],
			"queryMethod":    ServiceExampleFields["queryMethod"],
			"getEmptiesMsg":  ServiceExampleFields["getEmptiesMsg"],
		},
	})
	var Mutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"mutationMethod":     ServiceExampleMutationsFields["mutationMethod"],
			"getMutatuionMethod": ServiceExampleMutationsFields["getMutatuionMethod"],
		},
	})
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    Query,
		Mutation: Mutation,
	})
}
