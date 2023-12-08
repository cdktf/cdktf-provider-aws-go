// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxontapvolume


type FsxOntapVolumeSnaplockConfigurationRetentionPeriod struct {
	// default_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/fsx_ontap_volume#default_retention FsxOntapVolume#default_retention}
	DefaultRetention *FsxOntapVolumeSnaplockConfigurationRetentionPeriodDefaultRetention `field:"optional" json:"defaultRetention" yaml:"defaultRetention"`
	// maximum_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/fsx_ontap_volume#maximum_retention FsxOntapVolume#maximum_retention}
	MaximumRetention *FsxOntapVolumeSnaplockConfigurationRetentionPeriodMaximumRetention `field:"optional" json:"maximumRetention" yaml:"maximumRetention"`
	// minimum_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/fsx_ontap_volume#minimum_retention FsxOntapVolume#minimum_retention}
	MinimumRetention *FsxOntapVolumeSnaplockConfigurationRetentionPeriodMinimumRetention `field:"optional" json:"minimumRetention" yaml:"minimumRetention"`
}

