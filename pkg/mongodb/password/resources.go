package secret

import (
	"fmt"
	mongodbcontextstate "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/contextstate"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context) (*pulumi.Context, error) {
	randomPassword, err := random.NewRandomPassword(ctx, "generate-root-password", &random.RandomPasswordArgs{
		Length:     pulumi.Int(12),
		Special:    pulumi.Bool(true),
		Numeric:    pulumi.Bool(true),
		Upper:      pulumi.Bool(true),
		Lower:      pulumi.Bool(true),
		MinSpecial: pulumi.Int(3),
		MinNumeric: pulumi.Int(2),
		MinUpper:   pulumi.Int(2),
		MinLower:   pulumi.Int(2),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get random password value: %w", err)
	}

	var ctxState = ctx.Value(mongodbcontextstate.Key).(mongodbcontextstate.ContextState)

	addRandomPasswordToContext(&ctxState, randomPassword)
	ctx = ctx.WithValue(mongodbcontextstate.Key, ctxState)
	return ctx, nil
}

func addRandomPasswordToContext(existingConfig *mongodbcontextstate.ContextState, randomPassword *random.RandomPassword) {
	if existingConfig.Status.AddedResources == nil {
		existingConfig.Status.AddedResources = &mongodbcontextstate.AddedResources{
			RandomPassword: randomPassword,
		}
		return
	}
	existingConfig.Status.AddedResources.RandomPassword = randomPassword
}
