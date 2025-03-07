// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qbusinessapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QbusinessApplicationConfig struct {
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
	// The display name of the Amazon Q application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/qbusiness_application#display_name QbusinessApplication#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The Amazon Resource Name (ARN) of the IAM service role that provides permissions for the Amazon Q application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/qbusiness_application#iam_service_role_arn QbusinessApplication#iam_service_role_arn}
	IamServiceRoleArn *string `field:"required" json:"iamServiceRoleArn" yaml:"iamServiceRoleArn"`
	// ARN of the IAM Identity Center instance you are either creating for—or connecting to—your Amazon Q Business application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/qbusiness_application#identity_center_instance_arn QbusinessApplication#identity_center_instance_arn}
	IdentityCenterInstanceArn *string `field:"required" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
	// attachments_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/qbusiness_application#attachments_configuration QbusinessApplication#attachments_configuration}
	AttachmentsConfiguration interface{} `field:"optional" json:"attachmentsConfiguration" yaml:"attachmentsConfiguration"`
	// A description of the Amazon Q application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/qbusiness_application#description QbusinessApplication#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/qbusiness_application#encryption_configuration QbusinessApplication#encryption_configuration}
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/qbusiness_application#tags QbusinessApplication#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/qbusiness_application#timeouts QbusinessApplication#timeouts}
	Timeouts *QbusinessApplicationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

