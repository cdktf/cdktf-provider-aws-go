package route53record


type Route53RecordAlias struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/route53_record#evaluate_target_health Route53Record#evaluate_target_health}.
	EvaluateTargetHealth interface{} `field:"required" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/route53_record#name Route53Record#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/route53_record#zone_id Route53Record#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

