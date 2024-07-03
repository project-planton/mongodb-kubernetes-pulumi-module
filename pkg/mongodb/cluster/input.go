package mongodbkubernetes

import (
	mongodbcontextstate "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/contextstate"
	plantoncloudmongodbmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/mongodbkubernetes/model"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	resourceId       string
	namespace        *kubernetescorev1.Namespace
	kubeServiceName  string
	containerSpec    *plantoncloudmongodbmodel.MongodbKubernetesSpecContainerSpec
	customHelmValues map[string]string
	labels           map[string]string
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxState = ctx.Value(mongodbcontextstate.Key).(mongodbcontextstate.ContextState)

	return &input{
		resourceId:       ctxState.Spec.ResourceId,
		namespace:        ctxState.Status.AddedResources.Namespace,
		containerSpec:    ctxState.Spec.ContainerSpec,
		customHelmValues: ctxState.Spec.CustomHelmValues,
		labels:           ctxState.Spec.Labels,
		kubeServiceName:  ctxState.Spec.KubeServiceName,
	}
}
