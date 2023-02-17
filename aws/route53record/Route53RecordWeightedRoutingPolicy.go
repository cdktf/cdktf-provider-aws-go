package route53record


type Route53RecordWeightedRoutingPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/route53_record#weight Route53Record#weight}.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

