package outputs

import (
	mongodbcontextstate "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/contextstate"
	pulumicommonsloadbalancerservice "github.com/plantoncloud/pulumi-blueprint-commons/pkg/kubernetes/loadbalancer/service"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	rootUsername = "root"
	MongoDbPort  = 27017
)

type input struct {
	resourceId                    string
	resourceName                  string
	environmentName               string
	endpointDomainName            string
	namespaceName                 string
	externalLoadBalancerIpAddress string
	internalLoadBalancerIpAddress string
	internalHostname              string
	externalHostname              string
	kubeServiceName               string
	kubeLocalEndpoint             string
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxState = ctx.Value(mongodbcontextstate.Key).(mongodbcontextstate.ContextState)
	var externalLoadBalancerIpAddress = ""
	var internalLoadBalancerIpAddress = ""

	if ctxState.Status.AddedResources.LoadBalancerExternalService != nil {
		externalLoadBalancerIpAddress = pulumicommonsloadbalancerservice.GetIpAddress(ctxState.Status.AddedResources.LoadBalancerExternalService)
	}

	if ctxState.Status.AddedResources.LoadBalancerInternalService != nil {
		internalLoadBalancerIpAddress = pulumicommonsloadbalancerservice.GetIpAddress(ctxState.Status.AddedResources.LoadBalancerExternalService)
	}

	return &input{
		resourceId:                    ctxState.Spec.ResourceId,
		resourceName:                  ctxState.Spec.ResourceName,
		environmentName:               ctxState.Spec.EnvironmentInfo.EnvironmentName,
		endpointDomainName:            ctxState.Spec.EndpointDomainName,
		namespaceName:                 ctxState.Spec.NamespaceName,
		externalLoadBalancerIpAddress: externalLoadBalancerIpAddress,
		internalLoadBalancerIpAddress: internalLoadBalancerIpAddress,
		internalHostname:              ctxState.Spec.InternalHostname,
		externalHostname:              ctxState.Spec.ExternalHostname,
		kubeServiceName:               ctxState.Spec.KubeServiceName,
		kubeLocalEndpoint:             ctxState.Spec.KubeLocalEndpoint,
	}
}
