package pkg

var vars = struct {
	IstioIngressNamespace      string
	IstioIngressSelectorLabels map[string]string
	MongodbRootPasswordKey     string
	RootUsername               string
	MongoDbPort                int
}{
	IstioIngressNamespace: "istio-ingress",
	IstioIngressSelectorLabels: map[string]string{
		"app":   "istio-ingress",
		"istio": "ingress",
	},
	MongodbRootPasswordKey: "mongodb-root-password",
	RootUsername:           "root",
	MongoDbPort:            27017,
}
