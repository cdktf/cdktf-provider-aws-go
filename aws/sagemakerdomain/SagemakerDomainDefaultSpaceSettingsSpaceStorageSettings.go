// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultSpaceSettingsSpaceStorageSettings struct {
	// default_ebs_storage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/sagemaker_domain#default_ebs_storage_settings SagemakerDomain#default_ebs_storage_settings}
	DefaultEbsStorageSettings *SagemakerDomainDefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettings `field:"optional" json:"defaultEbsStorageSettings" yaml:"defaultEbsStorageSettings"`
}

