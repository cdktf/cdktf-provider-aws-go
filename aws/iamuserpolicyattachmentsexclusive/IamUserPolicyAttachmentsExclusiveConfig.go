// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamuserpolicyattachmentsexclusive

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamUserPolicyAttachmentsExclusiveConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/iam_user_policy_attachments_exclusive#policy_arns IamUserPolicyAttachmentsExclusive#policy_arns}.
	PolicyArns *[]*string `field:"required" json:"policyArns" yaml:"policyArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/iam_user_policy_attachments_exclusive#user_name IamUserPolicyAttachmentsExclusive#user_name}.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
}

