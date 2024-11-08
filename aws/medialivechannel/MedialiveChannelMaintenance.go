// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelMaintenance struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/medialive_channel#maintenance_day MedialiveChannel#maintenance_day}.
	MaintenanceDay *string `field:"required" json:"maintenanceDay" yaml:"maintenanceDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/medialive_channel#maintenance_start_time MedialiveChannel#maintenance_start_time}.
	MaintenanceStartTime *string `field:"required" json:"maintenanceStartTime" yaml:"maintenanceStartTime"`
}

