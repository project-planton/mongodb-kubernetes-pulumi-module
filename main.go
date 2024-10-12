package main

import (
	mongodbkubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/mongodbkubernetes/v1"
	"github.com/pkg/errors"
	"github.com/project-planton/mongodb-kubernetes-pulumi-module/pkg"
	"github.com/project-planton/pulumi-module-golang-commons/pkg/stackinput"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &mongodbkubernetesv1.MongodbKubernetesStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return pkg.Resources(ctx, stackInput)
	})
}
