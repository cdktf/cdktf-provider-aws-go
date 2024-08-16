// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsiampolicydocument

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsIamPolicyDocumentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/data-sources/iam_policy_document#id DataAwsIamPolicyDocument#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/data-sources/iam_policy_document#override_json DataAwsIamPolicyDocument#override_json}.
	OverrideJson *string `field:"optional" json:"overrideJson" yaml:"overrideJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/data-sources/iam_policy_document#override_policy_documents DataAwsIamPolicyDocument#override_policy_documents}.
	OverridePolicyDocuments *[]*string `field:"optional" json:"overridePolicyDocuments" yaml:"overridePolicyDocuments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/data-sources/iam_policy_document#policy_id DataAwsIamPolicyDocument#policy_id}.
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/data-sources/iam_policy_document#source_json DataAwsIamPolicyDocument#source_json}.
	SourceJson *string `field:"optional" json:"sourceJson" yaml:"sourceJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/data-sources/iam_policy_document#source_policy_documents DataAwsIamPolicyDocument#source_policy_documents}.
	SourcePolicyDocuments *[]*string `field:"optional" json:"sourcePolicyDocuments" yaml:"sourcePolicyDocuments"`
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/data-sources/iam_policy_document#statement DataAwsIamPolicyDocument#statement}
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/data-sources/iam_policy_document#version DataAwsIamPolicyDocument#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

