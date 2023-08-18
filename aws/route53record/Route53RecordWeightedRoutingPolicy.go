package route53record


type Route53RecordWeightedRoutingPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/route53_record#weight Route53Record#weight}.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

