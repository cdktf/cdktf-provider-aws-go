// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamrolepoliciesexclusive

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamRolePoliciesExclusiveConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/iam_role_policies_exclusive#policy_names IamRolePoliciesExclusive#policy_names}.
	PolicyNames *[]*string `field:"required" json:"policyNames" yaml:"policyNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/iam_role_policies_exclusive#role_name IamRolePoliciesExclusive#role_name}.
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
}

