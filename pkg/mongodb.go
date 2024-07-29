package pkg

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/kubernetes/containerresources"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/kubernetes/helm/mergemaps"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func mongodb(ctx *pulumi.Context, locals *Locals,
	createdNamespace *kubernetescorev1.Namespace, labels map[string]string) error {

	// https://github.com/bitnami/charts/blob/main/bitnami/mongodb/values.yaml
	var helmValues = pulumi.Map{
		"fullnameOverride":  pulumi.String(locals.KubeServiceName),
		"namespaceOverride": createdNamespace.Metadata.Name(),
		"resources":         containerresources.ConvertToPulumiMap(locals.MongodbKubernetes.Spec.Container.Resources),
		// todo: hard-coding this to 1 since we are only using `standalone` architecture,
		// need to revisit this to handle `replicaSet` architecture
		"replicaCount": pulumi.Int(1),
		"persistence": pulumi.Map{
			"enabled": pulumi.Bool(locals.MongodbKubernetes.Spec.Container.IsPersistenceEnabled),
			"size":    pulumi.String(locals.MongodbKubernetes.Spec.Container.DiskSize),
		},
		"podLabels":      pulumi.ToStringMap(labels),
		"commonLabels":   pulumi.ToStringMap(labels),
		"useStatefulSet": pulumi.Bool(true),
		"auth": pulumi.Map{
			"existingSecret": pulumi.String(locals.KubeServiceName),
		},
	}
	mergemaps.MergeMapToPulumiMap(helmValues, locals.MongodbKubernetes.Spec.HelmValues)

	// Deploying a Mongodb Helm chart from the Helm repository.
	_, err := helmv3.NewChart(ctx, locals.MongodbKubernetes.Metadata.Id, helmv3.ChartArgs{
		Chart:     pulumi.String("mongodb"),
		Version:   pulumi.String("15.1.4"), // Use the Helm chart version you want to install
		Namespace: pulumi.String(locals.Namespace),
		Values:    helmValues,
		//if you need to add the repository, you can specify `repo url`:
		// The URL for the Helm chart repository
		FetchArgs: helmv3.FetchArgs{
			Repo: pulumi.String("https://charts.bitnami.com/bitnami"),
		},
	}, pulumi.Timeouts(&pulumi.CustomTimeouts{Create: "2m", Update: "2m", Delete: "2m"}), pulumi.Parent(createdNamespace))

	if err != nil {
		return errors.Wrap(err, "failed to create mongodb resource")
	}
	return nil
}