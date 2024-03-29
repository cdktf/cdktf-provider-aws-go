// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsemrreleaselabels


type DataAwsEmrReleaseLabelsFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/data-sources/emr_release_labels#application DataAwsEmrReleaseLabels#application}.
	Application *string `field:"optional" json:"application" yaml:"application"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/data-sources/emr_release_labels#prefix DataAwsEmrReleaseLabels#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

