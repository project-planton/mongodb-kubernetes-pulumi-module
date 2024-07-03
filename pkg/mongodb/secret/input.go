package secret

import (
	mongodbcontextstate "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/contextstate"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	MongodbRootPasswordKey = "mongodb-root-password"
)

type input struct {
	namespaceName  string
	resourceName   string
	labels         map[string]string
	kubeProvider   *pulumikubernetes.Provider
	namespace      *kubernetescorev1.Namespace
	randomPassword *random.RandomPassword
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxConfig = ctx.Value(mongodbcontextstate.Key).(mongodbcontextstate.ContextState)

	return &input{
		namespaceName:  ctxConfig.Spec.NamespaceName,
		labels:         ctxConfig.Spec.Labels,
		kubeProvider:   ctxConfig.Spec.KubeProvider,
		resourceName:   ctxConfig.Spec.ResourceName,
		namespace:      ctxConfig.Status.AddedResources.Namespace,
		randomPassword: ctxConfig.Status.AddedResources.RandomPassword,
	}
}
