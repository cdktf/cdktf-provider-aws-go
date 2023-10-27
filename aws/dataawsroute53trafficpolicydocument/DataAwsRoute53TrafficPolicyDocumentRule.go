// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsroute53trafficpolicydocument


type DataAwsRoute53TrafficPolicyDocumentRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/data-sources/route53_traffic_policy_document#id DataAwsRoute53TrafficPolicyDocument#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// geo_proximity_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/data-sources/route53_traffic_policy_document#geo_proximity_location DataAwsRoute53TrafficPolicyDocument#geo_proximity_location}
	GeoProximityLocation interface{} `field:"optional" json:"geoProximityLocation" yaml:"geoProximityLocation"`
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/data-sources/route53_traffic_policy_document#items DataAwsRoute53TrafficPolicyDocument#items}
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/data-sources/route53_traffic_policy_document#location DataAwsRoute53TrafficPolicyDocument#location}
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// primary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/data-sources/route53_traffic_policy_document#primary DataAwsRoute53TrafficPolicyDocument#primary}
	Primary *DataAwsRoute53TrafficPolicyDocumentRulePrimary `field:"optional" json:"primary" yaml:"primary"`
	// region block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/data-sources/route53_traffic_policy_document#region DataAwsRoute53TrafficPolicyDocument#region}
	Region interface{} `field:"optional" json:"region" yaml:"region"`
	// secondary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/data-sources/route53_traffic_policy_document#secondary DataAwsRoute53TrafficPolicyDocument#secondary}
	Secondary *DataAwsRoute53TrafficPolicyDocumentRuleSecondary `field:"optional" json:"secondary" yaml:"secondary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/data-sources/route53_traffic_policy_document#type DataAwsRoute53TrafficPolicyDocument#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

