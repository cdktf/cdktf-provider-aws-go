// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssoadminaccountassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsoadminAccountAssignmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/ssoadmin_account_assignment#instance_arn SsoadminAccountAssignment#instance_arn}.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/ssoadmin_account_assignment#permission_set_arn SsoadminAccountAssignment#permission_set_arn}.
	PermissionSetArn *string `field:"required" json:"permissionSetArn" yaml:"permissionSetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/ssoadmin_account_assignment#principal_id SsoadminAccountAssignment#principal_id}.
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/ssoadmin_account_assignment#principal_type SsoadminAccountAssignment#principal_type}.
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/ssoadmin_account_assignment#target_id SsoadminAccountAssignment#target_id}.
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/ssoadmin_account_assignment#id SsoadminAccountAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/ssoadmin_account_assignment#target_type SsoadminAccountAssignment#target_type}.
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/ssoadmin_account_assignment#timeouts SsoadminAccountAssignment#timeouts}
	Timeouts *SsoadminAccountAssignmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

