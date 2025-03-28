// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controlmultiregionaccesspointpolicy


type S3ControlMultiRegionAccessPointPolicyDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/s3control_multi_region_access_point_policy#name S3ControlMultiRegionAccessPointPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/s3control_multi_region_access_point_policy#policy S3ControlMultiRegionAccessPointPolicy#policy}.
	Policy *string `field:"required" json:"policy" yaml:"policy"`
}

