// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmembership


type CleanroomsMembershipPaymentConfiguration struct {
	// query_compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/resources/cleanrooms_membership#query_compute CleanroomsMembership#query_compute}
	QueryCompute interface{} `field:"optional" json:"queryCompute" yaml:"queryCompute"`
}

