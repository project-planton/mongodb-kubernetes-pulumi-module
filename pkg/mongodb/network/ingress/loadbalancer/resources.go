package loadbalancer

import (
	"github.com/pkg/errors"
	mongodbcontextstate "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/contextstate"
	"github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/network/ingress/loadbalancer/gcp"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/cloudaccount/enums/kubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context) (newCtx *pulumi.Context, err error) {
	var ctxState = ctx.Value(mongodbcontextstate.Key).(mongodbcontextstate.ContextState)

	if ctxState.Spec.EnvironmentInfo.KubernetesProvider == kubernetesprovider.KubernetesProvider_gcp_gke {
		ctx, err = gcp.Resources(ctx)
		if err != nil {
			return ctx, errors.Wrap(err, "failed to create load balancer resources for gke cluster")
		}
	}
	return ctx, nil
}
