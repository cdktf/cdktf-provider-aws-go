// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace


type SagemakerSpaceSpaceSettingsSpaceStorageSettings struct {
	// ebs_storage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/sagemaker_space#ebs_storage_settings SagemakerSpace#ebs_storage_settings}
	EbsStorageSettings *SagemakerSpaceSpaceSettingsSpaceStorageSettingsEbsStorageSettings `field:"required" json:"ebsStorageSettings" yaml:"ebsStorageSettings"`
}

