package outputs

import (
	"fmt"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/english/enums/englishword"
	"github.com/plantoncloud/pulumi-blueprint-golang-commons/pkg/kubernetes/pulumikubernetesprovider"
	"github.com/plantoncloud/pulumi-blueprint-golang-commons/pkg/pulumi/pulumicustomoutput"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Export(ctx *pulumi.Context) error {
	var i = extractInput(ctx)
	var kubePortForwardCommand = getKubePortForwardCommand(i.namespaceName, i.resourceName)

	ctx.Export(GetExternalHostnameOutputName(), pulumi.String(i.externalHostname))
	ctx.Export(GetInternalHostnameOutputName(), pulumi.String(i.internalHostname))
	ctx.Export(GetKubeServiceNameOutputName(), pulumi.String(i.kubeServiceName))
	ctx.Export(GetKubeEndpointOutputName(), pulumi.String(i.kubeLocalEndpoint))
	ctx.Export(GetKubePortForwardCommandOutputName(), pulumi.String(kubePortForwardCommand))
	ctx.Export(GetExternalLoadBalancerIp(), pulumi.String(i.externalLoadBalancerIpAddress))
	ctx.Export(GetInternalLoadBalancerIp(), pulumi.String(i.internalLoadBalancerIpAddress))
	ctx.Export(GetNamespaceNameOutputName(), pulumi.String(i.namespaceName))
	ctx.Export(GetRootUsernameOutputName(), pulumi.String(rootUsername))
	ctx.Export(GetRootPasswordSecretOutputName(), pulumi.String(i.resourceName))
	return nil
}

func GetExternalHostnameOutputName() string {
	return pulumicustomoutput.Name("external-hostname")
}

func GetInternalHostnameOutputName() string {
	return pulumicustomoutput.Name("internal-hostname")
}

func GetKubeServiceNameOutputName() string {
	return pulumicustomoutput.Name("kubernetes-service-name")
}

func GetKubeEndpointOutputName() string {
	return pulumicustomoutput.Name("kubernetes-endpoint")
}

func GetKubePortForwardCommandOutputName() string {
	return pulumicustomoutput.Name("port-forward-command")
}

func GetExternalLoadBalancerIp() string {
	return pulumicustomoutput.Name("ingress-external-lb-ip")
}

func GetInternalLoadBalancerIp() string {
	return pulumicustomoutput.Name("ingress-internal-lb-ip")
}

func GetNamespaceNameOutputName() string {
	return pulumikubernetesprovider.PulumiOutputName(kubernetescorev1.Namespace{},
		englishword.EnglishWord_namespace.String())
}

// getKubePortForwardCommand returns kubectl port-forward command that can be used by developers.
// ex: "kubectl port-forward -n kubernetes_namespace  service/main-mongodb-kubernetes 8080:8080"
func getKubePortForwardCommand(namespaceName, kubeServiceName string) string {
	return fmt.Sprintf("kubectl port-forward -n %s service/%s %d:%d",
		namespaceName, kubeServiceName, MongoDbPort, MongoDbPort)
}

func GetRootUsernameOutputName() string {
	return pulumicustomoutput.Name("root-username")
}

func GetRootPasswordSecretOutputName() string {
	return pulumicustomoutput.Name("root-password-secret-name")
}
