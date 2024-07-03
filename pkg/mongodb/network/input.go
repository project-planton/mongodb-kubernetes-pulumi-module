package network

import (
	mongodbcontextstate "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/contextstate"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	isIngressEnabled   bool
	endpointDomainName string
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxState = ctx.Value(mongodbcontextstate.Key).(mongodbcontextstate.ContextState)

	return &input{
		isIngressEnabled:   ctxState.Spec.IsIngressEnabled,
		endpointDomainName: ctxState.Spec.EndpointDomainName,
	}
}
