// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package shielddrtaccessrolearnassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ShieldDrtAccessRoleArnAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/shield_drt_access_role_arn_association#role_arn ShieldDrtAccessRoleArnAssociation#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/shield_drt_access_role_arn_association#timeouts ShieldDrtAccessRoleArnAssociation#timeouts}
	Timeouts *ShieldDrtAccessRoleArnAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

