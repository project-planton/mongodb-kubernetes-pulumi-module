package gcp

import (
	"github.com/pkg/errors"
	mongodbclusterresources "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/cluster"
	mongodbcontextstate "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/contextstate"
	mongodbnamespaceresources "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/namespace"
	mongodbnetworkresources "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/network"
	mongodboutputs "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/outputs"
	mongodbpasswordresources "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/password"
	mongodbsecretresources "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/secret"
	mongodbkubernetesstackmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/mongodbkubernetes/stack/model"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceStack struct {
	WorkspaceDir     string
	Input            *mongodbkubernetesstackmodel.MongodbKubernetesStackInput
	KubernetesLabels map[string]string
}

func (resourceStack *ResourceStack) Resources(ctx *pulumi.Context) error {
	// https://artifacthub.io/packages/helm/bitnami/mongodb
	var ctxConfig, err = loadConfig(ctx, resourceStack)
	if err != nil {
		return errors.Wrap(err, "failed to initiate context config")
	}
	ctx = ctx.WithValue(mongodbcontextstate.Key, *ctxConfig)

	// Create the namespace resource
	ctx, err = mongodbnamespaceresources.Resources(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to create namespace resource")
	}

	// Create the random password resource
	ctx, err = mongodbpasswordresources.Resources(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to create random password resource")
	}

	// Create the secret resource for mongo db root password
	err = mongodbsecretresources.Resources(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to create secret resource")
	}

	// Deploying a Mongodb Helm chart from the Helm repository.
	err = mongodbclusterresources.Resources(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to create mongodb kubernetes")
	}

	ctx, err = mongodbnetworkresources.Resources(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to create mongodb network resources")
	}

	err = mongodboutputs.Export(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to export mongodb kubernetes outputs")
	}

	return nil
}
