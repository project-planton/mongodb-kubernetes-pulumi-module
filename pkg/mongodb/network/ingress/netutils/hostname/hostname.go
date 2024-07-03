package hostname

import (
	"fmt"
)

func GetInternalHostname(mongodbKubernetesId, environmentName, endpointDomainName string) string {
	return fmt.Sprintf("%s.%s-internal.%s", mongodbKubernetesId, environmentName, endpointDomainName)
}

func GetExternalHostname(mongodbKubernetesId, environmentName, endpointDomainName string) string {
	return fmt.Sprintf("%s.%s.%s", mongodbKubernetesId, environmentName, endpointDomainName)
}
