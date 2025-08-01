// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsquicksightuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsQuicksightUserConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/quicksight_user#user_name DataAwsQuicksightUser#user_name}.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/quicksight_user#aws_account_id DataAwsQuicksightUser#aws_account_id}.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/quicksight_user#id DataAwsQuicksightUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/quicksight_user#namespace DataAwsQuicksightUser#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/quicksight_user#region DataAwsQuicksightUser#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

