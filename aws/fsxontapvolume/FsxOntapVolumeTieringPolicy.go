// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxontapvolume


type FsxOntapVolumeTieringPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/fsx_ontap_volume#cooling_period FsxOntapVolume#cooling_period}.
	CoolingPeriod *float64 `field:"optional" json:"coolingPeriod" yaml:"coolingPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/fsx_ontap_volume#name FsxOntapVolume#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

