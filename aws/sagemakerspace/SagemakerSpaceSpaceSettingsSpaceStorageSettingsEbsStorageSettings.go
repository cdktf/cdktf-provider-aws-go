// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace


type SagemakerSpaceSpaceSettingsSpaceStorageSettingsEbsStorageSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/sagemaker_space#ebs_volume_size_in_gb SagemakerSpace#ebs_volume_size_in_gb}.
	EbsVolumeSizeInGb *float64 `field:"required" json:"ebsVolumeSizeInGb" yaml:"ebsVolumeSizeInGb"`
}

