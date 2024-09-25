// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controlmultiregionaccesspoint


type S3ControlMultiRegionAccessPointTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/s3control_multi_region_access_point#create S3ControlMultiRegionAccessPoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/s3control_multi_region_access_point#delete S3ControlMultiRegionAccessPoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

