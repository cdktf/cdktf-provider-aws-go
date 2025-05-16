// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmembership


type CleanroomsMembershipDefaultResultConfigurationOutputConfiguration struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/cleanrooms_membership#s3 CleanroomsMembership#s3}
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

