package ecsservice


type EcsServiceServiceConnectConfigurationServiceClientAlias struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/ecs_service#port EcsService#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/ecs_service#dns_name EcsService#dns_name}.
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
}

