// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmembership


type CleanroomsMembershipDefaultResultConfiguration struct {
	// output_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cleanrooms_membership#output_configuration CleanroomsMembership#output_configuration}
	OutputConfiguration interface{} `field:"optional" json:"outputConfiguration" yaml:"outputConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cleanrooms_membership#role_arn CleanroomsMembership#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

