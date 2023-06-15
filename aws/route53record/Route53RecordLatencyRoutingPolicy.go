package route53record


type Route53RecordLatencyRoutingPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/route53_record#region Route53Record#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
}

