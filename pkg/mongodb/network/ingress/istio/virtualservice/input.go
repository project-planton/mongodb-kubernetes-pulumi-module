package virtualservice

import (
	mongodbcontextstate "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/contextstate"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	resourceId        string
	namespace         *kubernetescorev1.Namespace
	hostNames         []string
	workspaceDir      string
	namespaceName     string
	kubeLocalEndpoint string
	kubeServiceName   string
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxState = ctx.Value(mongodbcontextstate.Key).(mongodbcontextstate.ContextState)

	return &input{
		resourceId:        ctxState.Spec.ResourceId,
		namespace:         ctxState.Status.AddedResources.Namespace,
		hostNames:         []string{ctxState.Spec.ExternalHostname},
		workspaceDir:      ctxState.Spec.WorkspaceDir,
		namespaceName:     ctxState.Spec.NamespaceName,
		kubeLocalEndpoint: ctxState.Spec.KubeLocalEndpoint,
		kubeServiceName:   ctxState.Spec.KubeServiceName,
	}
}
