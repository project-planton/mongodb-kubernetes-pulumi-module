package pkg

import (
	"github.com/pkg/errors"
	mongodbkubernetesmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/mongodbkubernetes"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/kubernetes/pulumikubernetesprovider"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	kubernetesmetav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceStack struct {
	Input  *mongodbkubernetesmodel.MongodbKubernetesStackInput
	Labels map[string]string
}

func (s *ResourceStack) Resources(ctx *pulumi.Context) error {
	locals := initializeLocals(ctx, s.Input)

	//create kubernetes-provider from the credential in the stack-input
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesClusterCredential(ctx,
		s.Input.KubernetesClusterCredential, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to create kubernetes provider")
	}

	//create namespace resource
	createdNamespace, err := kubernetescorev1.NewNamespace(ctx,
		locals.Namespace,
		&kubernetescorev1.NamespaceArgs{
			Metadata: kubernetesmetav1.ObjectMetaPtrInput(
				&kubernetesmetav1.ObjectMetaArgs{
					Name:   pulumi.String(locals.Namespace),
					Labels: pulumi.ToStringMap(s.Labels),
				}),
		},
		pulumi.Timeouts(&pulumi.CustomTimeouts{Create: "5s", Update: "5s", Delete: "5s"}),
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create %s namespace", locals.Namespace)
	}

	//create rootPasswordSecret resources
	if err := rootPasswordSecret(ctx, locals, createdNamespace, s.Labels); err != nil {
		return errors.Wrap(err, "failed to create root password secret resources")
	}

	//create rootPasswordSecret resources
	if err := mongodb(ctx, locals, createdNamespace, s.Labels); err != nil {
		return errors.Wrap(err, "failed to create mongodb helm-chart resources")
	}

	//create istio-ingress resources if ingress is enabled.
	if locals.MongodbKubernetes.Spec.Ingress.IsEnabled {
		if err := istioIngress(ctx, locals, createdNamespace, s.Labels); err != nil {
			return errors.Wrap(err, "failed to create istio ingress resources")
		}
	}

	return nil
}
