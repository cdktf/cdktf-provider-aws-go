package apprunnerservice


type ApprunnerServiceNetworkConfiguration struct {
	// egress_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service#egress_configuration ApprunnerService#egress_configuration}
	EgressConfiguration *ApprunnerServiceNetworkConfigurationEgressConfiguration `field:"optional" json:"egressConfiguration" yaml:"egressConfiguration"`
	// ingress_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service#ingress_configuration ApprunnerService#ingress_configuration}
	IngressConfiguration *ApprunnerServiceNetworkConfigurationIngressConfiguration `field:"optional" json:"ingressConfiguration" yaml:"ingressConfiguration"`
}

