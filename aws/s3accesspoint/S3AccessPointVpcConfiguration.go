// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3accesspoint


type S3AccessPointVpcConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_access_point#vpc_id S3AccessPoint#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

