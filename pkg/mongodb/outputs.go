package gcp

import (
	"context"
	"github.com/pkg/errors"
	"github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/outputs"
	"github.com/plantoncloud/pulumi-stack-runner-go-sdk/pkg/org"
	"github.com/plantoncloud/pulumi-stack-runner-go-sdk/pkg/stack/output/backend"

	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/enums/stackjoboperationtype"

	mongodbkubernetesmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/mongodbkubernetes/model"
	mongodbkubernetesstackmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/mongodbkubernetes/stack/model"
)

func Outputs(ctx context.Context, input *mongodbkubernetesstackmodel.MongodbKubernetesStackInput) (*mongodbkubernetesmodel.MongodbKubernetesStatusStackOutputs, error) {
	pulumiOrgName, err := org.GetOrgName()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get pulumi org name")
	}
	stackOutput, err := backend.StackOutput(pulumiOrgName, input.StackJob)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get stack output")
	}
	return OutputMapTransformer(stackOutput, input), nil
}

func OutputMapTransformer(stackOutput map[string]interface{}, input *mongodbkubernetesstackmodel.MongodbKubernetesStackInput) *mongodbkubernetesmodel.MongodbKubernetesStatusStackOutputs {
	if input.StackJob.Spec.OperationType != stackjoboperationtype.StackJobOperationType_apply || stackOutput == nil {
		return &mongodbkubernetesmodel.MongodbKubernetesStatusStackOutputs{}
	}
	return &mongodbkubernetesmodel.MongodbKubernetesStatusStackOutputs{
		Namespace:              backend.GetVal(stackOutput, outputs.GetNamespaceNameOutputName()),
		RootUsername:           backend.GetVal(stackOutput, outputs.GetRootUsernameOutputName()),
		RootPasswordSecretName: backend.GetVal(stackOutput, outputs.GetRootPasswordSecretOutputName()),
		Service:                backend.GetVal(stackOutput, outputs.GetKubeServiceNameOutputName()),
		PortForwardCommand:     backend.GetVal(stackOutput, outputs.GetKubePortForwardCommandOutputName()),
		KubeEndpoint:           backend.GetVal(stackOutput, outputs.GetKubeEndpointOutputName()),
		ExternalHostname:       backend.GetVal(stackOutput, outputs.GetExternalHostnameOutputName()),
		InternalHostname:       backend.GetVal(stackOutput, outputs.GetInternalHostnameOutputName()),
	}
}
