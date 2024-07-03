package network

import (
	"github.com/pkg/errors"
	mongodbingress "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/network/ingress"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context) (newCtx *pulumi.Context, err error) {
	i := extractInput(ctx)
	if !i.isIngressEnabled || i.endpointDomainName == "" {
		return ctx, nil
	}
	if ctx, err = mongodbingress.Resources(ctx); err != nil {
		return ctx, errors.Wrap(err, "failed to add gateway resources")
	}
	return ctx, nil
}
