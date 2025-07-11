// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/sagemaker_domain#default_ebs_volume_size_in_gb SagemakerDomain#default_ebs_volume_size_in_gb}.
	DefaultEbsVolumeSizeInGb *float64 `field:"required" json:"defaultEbsVolumeSizeInGb" yaml:"defaultEbsVolumeSizeInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/sagemaker_domain#maximum_ebs_volume_size_in_gb SagemakerDomain#maximum_ebs_volume_size_in_gb}.
	MaximumEbsVolumeSizeInGb *float64 `field:"required" json:"maximumEbsVolumeSizeInGb" yaml:"maximumEbsVolumeSizeInGb"`
}

