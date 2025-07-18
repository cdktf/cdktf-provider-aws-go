// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsroute53trafficpolicydocument


type DataAwsRoute53TrafficPolicyDocumentRuleLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/data-sources/route53_traffic_policy_document#continent DataAwsRoute53TrafficPolicyDocument#continent}.
	Continent *string `field:"optional" json:"continent" yaml:"continent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/data-sources/route53_traffic_policy_document#country DataAwsRoute53TrafficPolicyDocument#country}.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/data-sources/route53_traffic_policy_document#endpoint_reference DataAwsRoute53TrafficPolicyDocument#endpoint_reference}.
	EndpointReference *string `field:"optional" json:"endpointReference" yaml:"endpointReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/data-sources/route53_traffic_policy_document#evaluate_target_health DataAwsRoute53TrafficPolicyDocument#evaluate_target_health}.
	EvaluateTargetHealth interface{} `field:"optional" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/data-sources/route53_traffic_policy_document#health_check DataAwsRoute53TrafficPolicyDocument#health_check}.
	HealthCheck *string `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/data-sources/route53_traffic_policy_document#is_default DataAwsRoute53TrafficPolicyDocument#is_default}.
	IsDefault interface{} `field:"optional" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/data-sources/route53_traffic_policy_document#rule_reference DataAwsRoute53TrafficPolicyDocument#rule_reference}.
	RuleReference *string `field:"optional" json:"ruleReference" yaml:"ruleReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/data-sources/route53_traffic_policy_document#subdivision DataAwsRoute53TrafficPolicyDocument#subdivision}.
	Subdivision *string `field:"optional" json:"subdivision" yaml:"subdivision"`
}

