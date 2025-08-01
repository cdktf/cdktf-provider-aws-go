// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package macie2findingsfilter


type Macie2FindingsFilterFindingCriteria struct {
	// criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/macie2_findings_filter#criterion Macie2FindingsFilter#criterion}
	Criterion interface{} `field:"optional" json:"criterion" yaml:"criterion"`
}

