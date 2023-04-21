package route53record


type Route53RecordGeolocationRoutingPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/route53_record#continent Route53Record#continent}.
	Continent *string `field:"optional" json:"continent" yaml:"continent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/route53_record#country Route53Record#country}.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/route53_record#subdivision Route53Record#subdivision}.
	Subdivision *string `field:"optional" json:"subdivision" yaml:"subdivision"`
}

