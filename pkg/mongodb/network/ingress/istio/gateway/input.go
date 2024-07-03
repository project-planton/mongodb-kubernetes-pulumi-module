package gateway

import (
	mongodbcontextstate "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/contextstate"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	workspaceDir     string
	kubeProvider     *pulumikubernetes.Provider
	resourceId       string
	labels           map[string]string
	envDomainName    string
	externalHostname string
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxState = ctx.Value(mongodbcontextstate.Key).(mongodbcontextstate.ContextState)

	return &input{
		workspaceDir:     ctxState.Spec.WorkspaceDir,
		kubeProvider:     ctxState.Spec.KubeProvider,
		resourceId:       ctxState.Spec.ResourceId,
		labels:           ctxState.Spec.Labels,
		envDomainName:    ctxState.Spec.EnvDomainName,
		externalHostname: ctxState.Spec.ExternalHostname,
	}
}
