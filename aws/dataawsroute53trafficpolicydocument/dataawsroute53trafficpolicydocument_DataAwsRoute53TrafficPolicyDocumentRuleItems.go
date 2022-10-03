package dataawsroute53trafficpolicydocument


type DataAwsRoute53TrafficPolicyDocumentRuleItems struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/route53_traffic_policy_document#endpoint_reference DataAwsRoute53TrafficPolicyDocument#endpoint_reference}.
	EndpointReference *string `field:"optional" json:"endpointReference" yaml:"endpointReference"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/route53_traffic_policy_document#health_check DataAwsRoute53TrafficPolicyDocument#health_check}.
	HealthCheck *string `field:"optional" json:"healthCheck" yaml:"healthCheck"`
}

