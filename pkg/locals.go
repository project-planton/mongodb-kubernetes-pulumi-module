package pkg

import (
	mongodbkubernetesv1 "buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/apis/provider/kubernetes/mongodbkubernetes/v1"
	"fmt"
	"github.com/plantoncloud/mongodb-kubernetes-pulumi-module/pkg/outputs"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/kubernetes/kuberneteslabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	IngressExternalHostname  string
	IngressInternalHostname  string
	KubePortForwardCommand   string
	KubeServiceFqdn          string
	KubeServiceName          string
	KubernetesLabels         map[string]string
	MongodbKubernetes        *mongodbkubernetesv1.MongodbKubernetes
	Namespace                string
	MongodbPodSelectorLabels map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *mongodbkubernetesv1.MongodbKubernetesStackInput) *Locals {
	locals := &Locals{}
	//assign value for the locals variable to make it available across the project
	locals.MongodbKubernetes = stackInput.Target

	mongodbKubernetes := stackInput.Target

	locals.KubernetesLabels = map[string]string{
		kuberneteslabelkeys.Resource:     strconv.FormatBool(true),
		kuberneteslabelkeys.Organization: mongodbKubernetes.Spec.EnvironmentInfo.OrgId,
		kuberneteslabelkeys.Environment:  mongodbKubernetes.Spec.EnvironmentInfo.EnvId,
		kuberneteslabelkeys.ResourceKind: "mongodb_kubernetes",
		kuberneteslabelkeys.ResourceId:   mongodbKubernetes.Metadata.Id,
	}

	//decide on the namespace
	locals.Namespace = mongodbKubernetes.Metadata.Id

	ctx.Export(outputs.Namespace, pulumi.String(locals.Namespace))
	ctx.Export(outputs.RootUsername, pulumi.String(vars.RootUsername))
	ctx.Export(outputs.RootPasswordSecretName, pulumi.String(mongodbKubernetes.Metadata.Name))
	ctx.Export(outputs.RootPasswordSecretKey, pulumi.String(vars.MongodbRootPasswordKey))

	locals.KubeServiceName = mongodbKubernetes.Metadata.Name

	locals.MongodbPodSelectorLabels = map[string]string{
		"app.kubernetes.io/component": "mongodb",
		"app.kubernetes.io/instance":  mongodbKubernetes.Metadata.Id,
		"app.kubernetes.io/name":      "mongodb",
	}

	//export kubernetes service name
	ctx.Export(outputs.Service, pulumi.String(locals.KubeServiceName))

	locals.KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local", locals.KubeServiceName, locals.Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KubeEndpoint, pulumi.String(locals.KubeServiceFqdn))

	locals.KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		locals.Namespace, mongodbKubernetes.Metadata.Name)

	//export kube-port-forward command
	ctx.Export(outputs.KubePortForwardCommand, pulumi.String(locals.KubePortForwardCommand))

	if mongodbKubernetes.Spec.Ingress == nil ||
		!mongodbKubernetes.Spec.Ingress.IsEnabled ||
		mongodbKubernetes.Spec.Ingress.EndpointDomainName == "" {
		return locals
	}

	locals.IngressExternalHostname = fmt.Sprintf("%s.%s", mongodbKubernetes.Metadata.Id,
		mongodbKubernetes.Spec.Ingress.EndpointDomainName)

	locals.IngressInternalHostname = fmt.Sprintf("%s-internal.%s", mongodbKubernetes.Metadata.Id,
		mongodbKubernetes.Spec.Ingress.EndpointDomainName)

	//export ingress hostnames
	ctx.Export(outputs.IngressExternalHostname, pulumi.String(locals.IngressExternalHostname))
	ctx.Export(outputs.IngressInternalHostname, pulumi.String(locals.IngressInternalHostname))

	return locals
}
