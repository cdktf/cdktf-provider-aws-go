package ecsservice


type EcsServiceServiceConnectConfigurationService struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_service#port_name EcsService#port_name}.
	PortName *string `field:"required" json:"portName" yaml:"portName"`
	// client_alias block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_service#client_alias EcsService#client_alias}
	ClientAlias *EcsServiceServiceConnectConfigurationServiceClientAlias `field:"optional" json:"clientAlias" yaml:"clientAlias"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_service#discovery_name EcsService#discovery_name}.
	DiscoveryName *string `field:"optional" json:"discoveryName" yaml:"discoveryName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_service#ingress_port_override EcsService#ingress_port_override}.
	IngressPortOverride *float64 `field:"optional" json:"ingressPortOverride" yaml:"ingressPortOverride"`
}

