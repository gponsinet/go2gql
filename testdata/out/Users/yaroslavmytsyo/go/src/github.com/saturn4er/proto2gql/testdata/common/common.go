// This file was generated by github.com/EGT-Ukraine/go2gql. DO NOT EDIT IT
package common

import (
	context "context"
	debug "runtime/debug"

	errors "github.com/pkg/errors"
	graphql "github.com/saturn4er/graphql"

	scalars "github.com/EGT-Ukraine/go2gql/api/scalars"
	tracer "github.com/EGT-Ukraine/go2gql/api/tracer"
	common "github.com/EGT-Ukraine/go2gql/testdata/common"
)

// Enums
var CommonEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "CommonEnum",
	Description: "",
	Values: graphql.EnumValueConfigMap{
		"CommonEnumVal0": &graphql.EnumValueConfig{
			Value: 0,
		},
		"CommonEnumVal1": &graphql.EnumValueConfig{
			Value: 1,
		},
	},
})

// Input object
var CommonMessageInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CommonMessageInput",
	Fields: graphql.InputObjectConfigFieldMapThunk(func() graphql.InputObjectConfigFieldMap {
		return graphql.InputObjectConfigFieldMap{
			"scalar": &graphql.InputObjectFieldConfig{Type: scalars.GraphQLInt32Scalar},
		}
	}),
})

// Input objects resolvers
func ResolveCommonMessageInput(tr tracer.Tracer, ctx context.Context, i interface{}) (_ *common.CommonMessage, rerr error) {
	span := tr.CreateChildSpanFromContext(ctx, "ResolveCommonMessageInput")
	defer span.Finish()
	defer func() {
		if perr := recover(); perr != nil {
			span.SetTag("error", "true").SetTag("error_message", perr).SetTag("error_stack", string(debug.Stack()))
		}
		if rerr != nil {
			span.SetTag("error", "true").SetTag("error_message", rerr.Error())
		}
	}()
	if i == nil {
		return nil, nil
	}
	args := i.(map[string]interface{})
	_ = args
	var result = new(common.CommonMessage)
	if args["scalar"] != nil {
		result.Scalar = args["scalar"].(int32)
	}

	return result, nil
}

// Output objects
var CommonMessage = graphql.NewObject(graphql.ObjectConfig{
	Name:   "CommonMessage",
	Fields: graphql.Fields{},
})

func init() {
	CommonMessage.AddFieldConfig("scalar", &graphql.Field{
		Name: "scalar",
		Type: scalars.GraphQLInt32Scalar,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			switch src := p.Source.(type) {
			case *common.CommonMessage:
				if src == nil {
					return nil, nil
				}
				return src.Scalar, nil
			case common.CommonMessage:
				return src.Scalar, nil
			}
			return nil, errors.New("source of unknown type")
		},
	})
}

// Maps input objects
// Maps input objects resolvers
// Maps output objects
// Services
