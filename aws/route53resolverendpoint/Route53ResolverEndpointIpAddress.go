package route53resolverendpoint


type Route53ResolverEndpointIpAddress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/route53_resolver_endpoint#subnet_id Route53ResolverEndpoint#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/route53_resolver_endpoint#ip Route53ResolverEndpoint#ip}.
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
}

