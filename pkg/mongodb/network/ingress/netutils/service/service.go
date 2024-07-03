package service

import (
	"fmt"
	"github.com/plantoncloud-inc/go-commons/kubernetes/network/dns"
)

func GetKubeServiceNameFqdn(mongodbKubernetesName, namespace string) string {
	return fmt.Sprintf("%s.%s.%s", GetKubeServiceName(mongodbKubernetesName), namespace, dns.DefaultDomain)
}

func GetKubeServiceName(mongodbKubernetesName string) string {
	return fmt.Sprintf(mongodbKubernetesName)
}
