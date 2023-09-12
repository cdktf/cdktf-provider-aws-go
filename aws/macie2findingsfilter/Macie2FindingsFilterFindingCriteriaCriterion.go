// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package macie2findingsfilter


type Macie2FindingsFilterFindingCriteriaCriterion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/macie2_findings_filter#field Macie2FindingsFilter#field}.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/macie2_findings_filter#eq Macie2FindingsFilter#eq}.
	Eq *[]*string `field:"optional" json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/macie2_findings_filter#eq_exact_match Macie2FindingsFilter#eq_exact_match}.
	EqExactMatch *[]*string `field:"optional" json:"eqExactMatch" yaml:"eqExactMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/macie2_findings_filter#gt Macie2FindingsFilter#gt}.
	Gt *string `field:"optional" json:"gt" yaml:"gt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/macie2_findings_filter#gte Macie2FindingsFilter#gte}.
	Gte *string `field:"optional" json:"gte" yaml:"gte"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/macie2_findings_filter#lt Macie2FindingsFilter#lt}.
	Lt *string `field:"optional" json:"lt" yaml:"lt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/macie2_findings_filter#lte Macie2FindingsFilter#lte}.
	Lte *string `field:"optional" json:"lte" yaml:"lte"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/macie2_findings_filter#neq Macie2FindingsFilter#neq}.
	Neq *[]*string `field:"optional" json:"neq" yaml:"neq"`
}

