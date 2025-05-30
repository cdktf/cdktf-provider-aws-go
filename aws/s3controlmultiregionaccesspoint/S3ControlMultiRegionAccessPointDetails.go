// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controlmultiregionaccesspoint


type S3ControlMultiRegionAccessPointDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/s3control_multi_region_access_point#name S3ControlMultiRegionAccessPoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// region block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/s3control_multi_region_access_point#region S3ControlMultiRegionAccessPoint#region}
	Region interface{} `field:"required" json:"region" yaml:"region"`
	// public_access_block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/s3control_multi_region_access_point#public_access_block S3ControlMultiRegionAccessPoint#public_access_block}
	PublicAccessBlock *S3ControlMultiRegionAccessPointDetailsPublicAccessBlock `field:"optional" json:"publicAccessBlock" yaml:"publicAccessBlock"`
}

