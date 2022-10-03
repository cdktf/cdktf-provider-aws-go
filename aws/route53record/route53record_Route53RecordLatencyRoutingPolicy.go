package route53record


type Route53RecordLatencyRoutingPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/route53_record#region Route53Record#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
}

