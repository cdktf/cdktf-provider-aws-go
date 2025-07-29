// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxontapvolume


type FsxOntapVolumeAggregateConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/fsx_ontap_volume#aggregates FsxOntapVolume#aggregates}.
	Aggregates *[]*string `field:"optional" json:"aggregates" yaml:"aggregates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/fsx_ontap_volume#constituents_per_aggregate FsxOntapVolume#constituents_per_aggregate}.
	ConstituentsPerAggregate *float64 `field:"optional" json:"constituentsPerAggregate" yaml:"constituentsPerAggregate"`
}

