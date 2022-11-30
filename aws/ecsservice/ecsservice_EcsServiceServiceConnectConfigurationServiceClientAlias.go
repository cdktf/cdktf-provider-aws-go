package ecsservice


type EcsServiceServiceConnectConfigurationServiceClientAlias struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_service#port EcsService#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ecs_service#dns_name EcsService#dns_name}.
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
}

