package gcp

import (
	mongodbcontextstate "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/contextstate"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	resourceId         string
	resourceName       string
	namespace          *kubernetescorev1.Namespace
	externalHostname   string
	internalHostname   string
	endpointDomainName string
	serviceName        string
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxState = ctx.Value(mongodbcontextstate.Key).(mongodbcontextstate.ContextState)

	return &input{
		resourceId:         ctxState.Spec.ResourceId,
		resourceName:       ctxState.Spec.ResourceName,
		namespace:          ctxState.Status.AddedResources.Namespace,
		externalHostname:   ctxState.Spec.ExternalHostname,
		internalHostname:   ctxState.Spec.InternalHostname,
		endpointDomainName: ctxState.Spec.EndpointDomainName,
		serviceName:        ctxState.Spec.KubeServiceName,
	}
}
