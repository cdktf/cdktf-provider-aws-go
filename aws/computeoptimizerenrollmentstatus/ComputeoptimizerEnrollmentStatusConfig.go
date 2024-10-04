// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeoptimizerenrollmentstatus

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeoptimizerEnrollmentStatusConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/computeoptimizer_enrollment_status#status ComputeoptimizerEnrollmentStatus#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/computeoptimizer_enrollment_status#include_member_accounts ComputeoptimizerEnrollmentStatus#include_member_accounts}.
	IncludeMemberAccounts interface{} `field:"optional" json:"includeMemberAccounts" yaml:"includeMemberAccounts"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/computeoptimizer_enrollment_status#timeouts ComputeoptimizerEnrollmentStatus#timeouts}
	Timeouts *ComputeoptimizerEnrollmentStatusTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

