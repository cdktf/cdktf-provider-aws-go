// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawscloudwatchlogdataprotectionpolicydocument

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsCloudwatchLogDataProtectionPolicyDocumentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/data-sources/cloudwatch_log_data_protection_policy_document#name DataAwsCloudwatchLogDataProtectionPolicyDocument#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/data-sources/cloudwatch_log_data_protection_policy_document#statement DataAwsCloudwatchLogDataProtectionPolicyDocument#statement}
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/data-sources/cloudwatch_log_data_protection_policy_document#description DataAwsCloudwatchLogDataProtectionPolicyDocument#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/data-sources/cloudwatch_log_data_protection_policy_document#id DataAwsCloudwatchLogDataProtectionPolicyDocument#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/data-sources/cloudwatch_log_data_protection_policy_document#version DataAwsCloudwatchLogDataProtectionPolicyDocument#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

