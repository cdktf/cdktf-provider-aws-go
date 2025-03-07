// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssoadminpermissionsetinlinepolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsoadminPermissionSetInlinePolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/ssoadmin_permission_set_inline_policy#inline_policy SsoadminPermissionSetInlinePolicy#inline_policy}.
	InlinePolicy *string `field:"required" json:"inlinePolicy" yaml:"inlinePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/ssoadmin_permission_set_inline_policy#instance_arn SsoadminPermissionSetInlinePolicy#instance_arn}.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/ssoadmin_permission_set_inline_policy#permission_set_arn SsoadminPermissionSetInlinePolicy#permission_set_arn}.
	PermissionSetArn *string `field:"required" json:"permissionSetArn" yaml:"permissionSetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/ssoadmin_permission_set_inline_policy#id SsoadminPermissionSetInlinePolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/ssoadmin_permission_set_inline_policy#timeouts SsoadminPermissionSetInlinePolicy#timeouts}
	Timeouts *SsoadminPermissionSetInlinePolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

