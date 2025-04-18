// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmcontactsrotation


type SsmcontactsRotationRecurrenceShiftCoverages struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/ssmcontacts_rotation#map_block_key SsmcontactsRotation#map_block_key}.
	MapBlockKey *string `field:"required" json:"mapBlockKey" yaml:"mapBlockKey"`
	// coverage_times block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/ssmcontacts_rotation#coverage_times SsmcontactsRotation#coverage_times}
	CoverageTimes interface{} `field:"optional" json:"coverageTimes" yaml:"coverageTimes"`
}

