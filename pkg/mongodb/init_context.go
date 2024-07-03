package gcp

import (
	"github.com/pkg/errors"
	environmentblueprinthostnames "github.com/plantoncloud/environment-pulumi-blueprint/pkg/gcpgke/endpointdomains/hostnames"
	mongodbcontextstate "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/contextstate"
	mongodbnetutilshostname "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/network/ingress/netutils/hostname"
	mongodbnetutilsservice "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/network/ingress/netutils/service"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/enums/kubernetesworkloadingresstype"
	"github.com/plantoncloud/pulumi-blueprint-golang-commons/pkg/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func loadConfig(ctx *pulumi.Context, resourceStack *ResourceStack) (*mongodbcontextstate.ContextState, error) {

	kubernetesProvider, err := pulumikubernetesprovider.GetWithStackCredentials(ctx, resourceStack.Input.CredentialsInput)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup kubernetes provider")
	}

	var resourceId = resourceStack.Input.ResourceInput.Metadata.Id
	var resourceName = resourceStack.Input.ResourceInput.Metadata.Name
	var environmentInfo = resourceStack.Input.ResourceInput.Spec.EnvironmentInfo
	var isIngressEnabled = false

	if resourceStack.Input.ResourceInput.Spec.Ingress != nil {
		isIngressEnabled = resourceStack.Input.ResourceInput.Spec.Ingress.IsEnabled
	}

	var endpointDomainName = ""
	var envDomainName = ""
	var ingressType = kubernetesworkloadingresstype.KubernetesWorkloadIngressType_unspecified
	var internalHostname = ""
	var externalHostname = ""

	if isIngressEnabled {
		endpointDomainName = resourceStack.Input.ResourceInput.Spec.Ingress.EndpointDomainName
		envDomainName = environmentblueprinthostnames.GetExternalEnvHostname(environmentInfo.EnvironmentName, endpointDomainName)
		ingressType = resourceStack.Input.ResourceInput.Spec.Ingress.IngressType

		internalHostname = mongodbnetutilshostname.GetInternalHostname(resourceId, environmentInfo.EnvironmentName, endpointDomainName)
		externalHostname = mongodbnetutilshostname.GetExternalHostname(resourceId, environmentInfo.EnvironmentName, endpointDomainName)
	}

	return &mongodbcontextstate.ContextState{
		Spec: &mongodbcontextstate.Spec{
			KubeProvider:       kubernetesProvider,
			ResourceId:         resourceId,
			ResourceName:       resourceName,
			Labels:             resourceStack.KubernetesLabels,
			WorkspaceDir:       resourceStack.WorkspaceDir,
			NamespaceName:      resourceId,
			EnvironmentInfo:    resourceStack.Input.ResourceInput.Spec.EnvironmentInfo,
			IsIngressEnabled:   isIngressEnabled,
			IngressType:        ingressType,
			EndpointDomainName: endpointDomainName,
			EnvDomainName:      envDomainName,
			ContainerSpec:      resourceStack.Input.ResourceInput.Spec.Container,
			CustomHelmValues:   resourceStack.Input.ResourceInput.Spec.HelmValues,
			InternalHostname:   internalHostname,
			ExternalHostname:   externalHostname,
			KubeServiceName:    mongodbnetutilsservice.GetKubeServiceName(resourceName),
			KubeLocalEndpoint:  mongodbnetutilsservice.GetKubeServiceNameFqdn(resourceName, resourceId),
		},
		Status: &mongodbcontextstate.Status{},
	}, nil
}
