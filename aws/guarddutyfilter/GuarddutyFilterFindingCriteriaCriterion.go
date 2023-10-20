// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guarddutyfilter


type GuarddutyFilterFindingCriteriaCriterion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/guardduty_filter#field GuarddutyFilter#field}.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/guardduty_filter#equals GuarddutyFilter#equals}.
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/guardduty_filter#greater_than GuarddutyFilter#greater_than}.
	GreaterThan *string `field:"optional" json:"greaterThan" yaml:"greaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/guardduty_filter#greater_than_or_equal GuarddutyFilter#greater_than_or_equal}.
	GreaterThanOrEqual *string `field:"optional" json:"greaterThanOrEqual" yaml:"greaterThanOrEqual"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/guardduty_filter#less_than GuarddutyFilter#less_than}.
	LessThan *string `field:"optional" json:"lessThan" yaml:"lessThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/guardduty_filter#less_than_or_equal GuarddutyFilter#less_than_or_equal}.
	LessThanOrEqual *string `field:"optional" json:"lessThanOrEqual" yaml:"lessThanOrEqual"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/resources/guardduty_filter#not_equals GuarddutyFilter#not_equals}.
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
}

