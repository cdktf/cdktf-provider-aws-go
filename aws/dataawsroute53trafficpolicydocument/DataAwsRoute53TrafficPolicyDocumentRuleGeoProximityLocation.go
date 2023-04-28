package dataawsroute53trafficpolicydocument


type DataAwsRoute53TrafficPolicyDocumentRuleGeoProximityLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/route53_traffic_policy_document#bias DataAwsRoute53TrafficPolicyDocument#bias}.
	Bias *string `field:"optional" json:"bias" yaml:"bias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/route53_traffic_policy_document#endpoint_reference DataAwsRoute53TrafficPolicyDocument#endpoint_reference}.
	EndpointReference *string `field:"optional" json:"endpointReference" yaml:"endpointReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/route53_traffic_policy_document#evaluate_target_health DataAwsRoute53TrafficPolicyDocument#evaluate_target_health}.
	EvaluateTargetHealth interface{} `field:"optional" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/route53_traffic_policy_document#health_check DataAwsRoute53TrafficPolicyDocument#health_check}.
	HealthCheck *string `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/route53_traffic_policy_document#latitude DataAwsRoute53TrafficPolicyDocument#latitude}.
	Latitude *string `field:"optional" json:"latitude" yaml:"latitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/route53_traffic_policy_document#longitude DataAwsRoute53TrafficPolicyDocument#longitude}.
	Longitude *string `field:"optional" json:"longitude" yaml:"longitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/route53_traffic_policy_document#region DataAwsRoute53TrafficPolicyDocument#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/route53_traffic_policy_document#rule_reference DataAwsRoute53TrafficPolicyDocument#rule_reference}.
	RuleReference *string `field:"optional" json:"ruleReference" yaml:"ruleReference"`
}

