package pkg

import (
	"encoding/base64"
	"fmt"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func rootPasswordSecret(ctx *pulumi.Context, locals *Locals,
	createdNamespace *kubernetescorev1.Namespace, labels map[string]string) error {
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
		return fmt.Errorf("failed to generate random password value: %w", err)
	}

	// Encode the password in Base64
	base64Password := randomPassword.Result.ApplyT(func(p string) (string, error) {
		return base64.StdEncoding.EncodeToString([]byte(p)), nil
	}).(pulumi.StringOutput)

	// Create or update the secret
	_, err = kubernetescorev1.NewSecret(ctx, locals.MongodbKubernetes.Metadata.Name, &kubernetescorev1.SecretArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name:      pulumi.String(locals.MongodbKubernetes.Metadata.Name),
			Namespace: createdNamespace.Metadata.Name(),
		},
		Data: pulumi.StringMap{
			vars.MongodbRootPasswordKey: base64Password,
		},
	}, pulumi.Parent(createdNamespace), pulumi.Parent(randomPassword))

	return nil

}
