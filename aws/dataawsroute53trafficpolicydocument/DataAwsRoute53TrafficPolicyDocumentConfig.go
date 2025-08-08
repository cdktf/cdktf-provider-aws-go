// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsroute53trafficpolicydocument

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsRoute53TrafficPolicyDocumentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/route53_traffic_policy_document#endpoint DataAwsRoute53TrafficPolicyDocument#endpoint}
	Endpoint interface{} `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/route53_traffic_policy_document#id DataAwsRoute53TrafficPolicyDocument#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/route53_traffic_policy_document#record_type DataAwsRoute53TrafficPolicyDocument#record_type}.
	RecordType *string `field:"optional" json:"recordType" yaml:"recordType"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/route53_traffic_policy_document#rule DataAwsRoute53TrafficPolicyDocument#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/route53_traffic_policy_document#start_endpoint DataAwsRoute53TrafficPolicyDocument#start_endpoint}.
	StartEndpoint *string `field:"optional" json:"startEndpoint" yaml:"startEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/route53_traffic_policy_document#start_rule DataAwsRoute53TrafficPolicyDocument#start_rule}.
	StartRule *string `field:"optional" json:"startRule" yaml:"startRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/data-sources/route53_traffic_policy_document#version DataAwsRoute53TrafficPolicyDocument#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

