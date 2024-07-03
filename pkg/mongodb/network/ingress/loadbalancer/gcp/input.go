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
	var ctxConfig = ctx.Value(mongodbcontextstate.Key).(mongodbcontextstate.ContextState)

	return &input{
		resourceId:         ctxConfig.Spec.ResourceId,
		resourceName:       ctxConfig.Spec.ResourceName,
		namespace:          ctxConfig.Status.AddedResources.Namespace,
		externalHostname:   ctxConfig.Spec.ExternalHostname,
		internalHostname:   ctxConfig.Spec.InternalHostname,
		endpointDomainName: ctxConfig.Spec.EndpointDomainName,
		serviceName:        ctxConfig.Spec.KubeServiceName,
	}
}
