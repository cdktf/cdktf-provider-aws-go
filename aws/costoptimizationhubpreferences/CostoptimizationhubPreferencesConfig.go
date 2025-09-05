// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package costoptimizationhubpreferences

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CostoptimizationhubPreferencesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/costoptimizationhub_preferences#member_account_discount_visibility CostoptimizationhubPreferences#member_account_discount_visibility}.
	MemberAccountDiscountVisibility *string `field:"optional" json:"memberAccountDiscountVisibility" yaml:"memberAccountDiscountVisibility"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/costoptimizationhub_preferences#savings_estimation_mode CostoptimizationhubPreferences#savings_estimation_mode}.
	SavingsEstimationMode *string `field:"optional" json:"savingsEstimationMode" yaml:"savingsEstimationMode"`
}

